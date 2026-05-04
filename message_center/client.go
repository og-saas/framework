package message_center

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/og-saas/framework/utils/httpc"
	"github.com/zeromicro/go-zero/core/jsonx"
)

// Client 消息中心客户端
type Client struct {
	config Config
	client *http.Client
}

// NewClient 创建消息中心客户端
func NewClient(config Config) (*Client, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}

	return &Client{
		config: config,
		client: http.DefaultClient,
	}, nil
}

// generateSign 生成签名
func (c *Client) generateSign(method, path string, query map[string]string, timestamp int64, nonce string) string {
	// 构建 canonicalString: method + '\n' + path + '\n' + sortedQuery + '\n' + timestamp + '\n' + nonce
	sortedQuery := buildSortedQuery(query)
	canonicalString := fmt.Sprintf("%s\n%s\n%s\n%d\n%s", method, path, sortedQuery, timestamp, nonce)

	// 1. 先对 appSecret 做 SHA256 得到密钥
	secretHash := sha256.Sum256([]byte(c.config.AppSecret))

	// 2. 用派生密钥做 HMAC-SHA256
	h := hmac.New(sha256.New, secretHash[:])
	h.Write([]byte(canonicalString))
	return hex.EncodeToString(h.Sum(nil))
}

// buildSortedQuery 构建排序后的查询字符串
func buildSortedQuery(query map[string]string) string {
	if len(query) == 0 {
		return ""
	}

	keys := make([]string, 0, len(query))
	for k := range query {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for i, k := range keys {
		if i > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(query[k])
	}
	return buf.String()
}

// buildHeaders 构建请求头
func (c *Client) buildHeaders(method, path string) map[string]string {
	timestamp := time.Now().UnixMilli()
	nonce := generateNonce()

	sign := c.generateSign(method, path, nil, timestamp, nonce)
	return map[string]string{
		HeaderAppKey:    c.config.AppKey,
		HeaderTimestamp: strconv.FormatInt(timestamp, 10),
		HeaderNonce:     nonce,
		HeaderSign:      sign,
	}
}

// Otp 获取连接凭证
func (c *Client) Otp(ctx context.Context, req OtpReq) (*OtpResp, error) {
	// 构建内部请求，自动填充 AppKey
	internalReq := otpReqInternal{
		AppKey:   c.config.AppKey,
		ClientId: req.ClientId,
		Topics:   req.Topics,
	}

	return doRequestAndParse[OtpResp](c, ctx, OtpURL, http.MethodPost, internalReq)
}

// NormalSend 发送即时消息
func (c *Client) NormalSend(ctx context.Context, req SendMessageReq) (*SendMessageResp, error) {
	// 构建内部请求，自动填充 AppKey
	internalReq := sendMessageReqInternal{
		AppKey:             c.config.AppKey,
		Topic:              req.Topic,
		Content:            req.Content,
		Qos:                req.Qos,
		Retain:             req.Retain,
		RetainTimeDuration: req.RetainTimeDuration,
	}

	return doRequestAndParse[SendMessageResp](c, ctx, NormalSendURL, http.MethodPost, internalReq)
}

// TimerSend 发送定时消息
func (c *Client) TimerSend(ctx context.Context, req SendTimerMessageReq) (*SendTimerMessageResp, error) {
	// 构建内部请求，自动填充 AppKey
	internalReq := sendTimerMessageReqInternal{
		AppKey:             c.config.AppKey,
		Topic:              req.Topic,
		Content:            req.Content,
		Qos:                req.Qos,
		Retain:             req.Retain,
		RetainTimeDuration: req.RetainTimeDuration,
		SendTime:           req.SendTime,
	}

	return doRequestAndParse[SendTimerMessageResp](c, ctx, TimerSendURL, http.MethodPost, internalReq)
}

// CancelTimer 取消定时消息
func (c *Client) CancelTimer(ctx context.Context, req CancelTimerMessageReq) (*CancelTimerMessageResp, error) {
	// 构建内部请求，自动填充 AppKey
	internalReq := cancelTimerMessageReqInternal{
		AppKey:    c.config.AppKey,
		MessageId: req.MessageId,
	}

	return doRequestAndParse[CancelTimerMessageResp](c, ctx, TimerCancelURL, http.MethodPost, internalReq)
}

// QueryHistory 查询历史消息
func (c *Client) QueryHistory(ctx context.Context, req QueryHistoryReq) (*QueryHistoryResp, error) {
	return doRequestAndParse[QueryHistoryResp](c, ctx, HistoryQueryURL, http.MethodPost, req)
}

// doRequest 发送请求
func (c *Client) doRequest(ctx context.Context, path, method string, body interface{}) ([]byte, error) {
	// 检查 context 是否已有 deadline
	if _, ok := ctx.Deadline(); !ok {
		// 没有则使用默认超时作为兜底
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, c.config.getTimeout())
		defer cancel()
	}

	apiUrl := fmt.Sprintf("%s%s", c.config.HttpURL, path)
	headers := c.buildHeaders(method, path)

	client := httpc.Do(ctx).SetHeaders(headers)
	if body != nil {
		client = client.SetBody(body) // resty 会自动序列化为 JSON
	}

	var resp *resty.Response
	var err error

	switch method {
	case http.MethodGet:
		resp, err = client.Get(apiUrl)
	case http.MethodPost:
		resp, err = client.Post(apiUrl)
	default:
		return nil, fmt.Errorf("unsupported http method: %s", method)
	}

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode(), resp.String())
	}

	return resp.Body(), nil
}

// parseResponse 解析响应
func (c *Client) parseResponse(ctx context.Context, respBody []byte, result interface{}) error {
	if err := jsonx.Unmarshal(respBody, result); err != nil {
		return fmt.Errorf("response unmarshal error: %w", err)
	}
	return nil
}

// checkResponse 检查响应状态
func checkResponse(resp *CommonResp[any]) error {
	if resp.Code != 200 {
		return fmt.Errorf("response error: code=%d, msg=%s", resp.Code, resp.Msg)
	}
	return nil
}

// doRequestAndParse 发送请求并解析响应（泛型方法）
func doRequestAndParse[T any](c *Client, ctx context.Context, path, method string, body interface{}) (*T, error) {
	respBody, err := c.doRequest(ctx, path, method, body)
	if err != nil {
		return nil, err
	}

	var result CommonResp[T]
	if err := jsonx.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("response unmarshal error: %w", err)
	}

	if result.Code != 200 {
		return nil, fmt.Errorf("response error: code=%d, msg=%s", result.Code, result.Msg)
	}

	return &result.Data, nil
}

// generateNonce 生成随机字符串
func generateNonce() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		// 降级到时间戳方案
		return strconv.FormatInt(time.Now().UnixNano(), 36)
	}
	return hex.EncodeToString(b)
}
