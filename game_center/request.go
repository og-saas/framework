package game_center

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/og-saas/framework/utils/httpc"
	"github.com/og-saas/framework/utils/xerr"
)

func postRequest(ctx context.Context, config CenterConfig, reqURL, currency string, body any) (*resty.Response, error) {
	resp, err := httpc.Do(ctx).
		SetBasicAuth(config.GetCurrencyConf(currency).Username, config.GetCurrencyConf(currency).Password).
		SetBody(body).
		Post(config.RequestURL + reqURL)
	if err != nil {
		return nil, wrapNetworkError(err)
	}

	if err := checkIntercepted(resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New("POST request failed with status " + resp.Status() + ": " + resp.String())
	}
	return resp, nil
}

func getRequest(ctx context.Context, config CenterConfig, reqURL, currency string, params url.Values) (*resty.Response, error) {
	resp, err := httpc.Do(ctx).
		SetBasicAuth(config.GetCurrencyConf(currency).Username, config.GetCurrencyConf(currency).Password).
		SetQueryParamsFromValues(params).
		Get(config.RequestURL + reqURL)
	if err != nil {
		return nil, wrapNetworkError(err)
	}

	if err := checkIntercepted(resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New("GET request failed with status " + resp.Status() + ": " + resp.String())
	}
	return resp, nil
}

// wrapNetworkError 判断网络层错误，能确定没到中台的返回 ErrCodeGamePlatformUnreachable
func wrapNetworkError(err error) error {
	if err == nil {
		return nil
	}

	// DNS 解析失败
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		return xerr.NewError(xerr.ErrCodeGamePlatformUnreachable, err.Error())
	}

	// 连接被拒绝（端口没监听/服务没启动）
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		if errors.Is(opErr.Err, net.ErrClosed) || errors.Is(opErr, net.ErrClosed) {
			return xerr.NewError(xerr.ErrCodeGamePlatformUnreachable, err.Error())
		}
		// connection refused
		if strings.Contains(opErr.Error(), "connection refused") {
			return xerr.NewError(xerr.ErrCodeGamePlatformUnreachable, err.Error())
		}
	}

	// 连接超时（connect timeout，不是 read timeout）
	if strings.Contains(err.Error(), "connection timed out") ||
		strings.Contains(err.Error(), "no route to host") ||
		strings.Contains(err.Error(), "network is unreachable") ||
		strings.Contains(err.Error(), "TLS handshake") {
		return xerr.NewError(xerr.ErrCodeGamePlatformUnreachable, err.Error())
	}

	// 其他网络错误（read timeout、context deadline 等）不能确定没到中台，原样返回
	return err
}

// checkIntercepted 检查响应是否是被 WAF/CDN/反代拦截返回的 HTML 页面
func checkIntercepted(resp *resty.Response) error {
	contentType := resp.Header().Get("Content-Type")
	body := resp.String()

	// Content-Type 是 text/html，说明不是中台的 JSON 响应
	if strings.Contains(strings.ToLower(contentType), "text/html") {
		return xerr.NewError(xerr.ErrCodeGamePlatformUnreachable,
			"request intercepted, received HTML response: "+truncate(body, 200))
	}

	// 兜底：body 以 HTML 标签开头，即使 Content-Type 没标 text/html
	trimmed := strings.TrimSpace(body)
	if strings.HasPrefix(trimmed, "<!") || strings.HasPrefix(trimmed, "<html") || strings.HasPrefix(trimmed, "<HTML") {
		return xerr.NewError(xerr.ErrCodeGamePlatformUnreachable,
			"request intercepted, received HTML response: "+truncate(body, 200))
	}

	return nil
}

func truncate(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen] + "..."
	}
	return s
}
