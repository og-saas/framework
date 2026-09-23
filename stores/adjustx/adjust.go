package adjustx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/fatih/structs"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest/httpc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Sender struct {
	ctx  context.Context
	url  string
	Auth string
}

func NewAdjustSender(ctx context.Context, eventConfig Config) *Sender {
	return &Sender{
		ctx:  ctx,
		url:  eventConfig.Url,
		Auth: eventConfig.Auth,
	}
}

func (sender *Sender) GetEnv(env string) string {
	if env == service.DevMode || env == service.TestMode {
		return "sandbox"
	}
	return "production"
}

func (sender *Sender) GetAuth() string {
	return fmt.Sprintf("Bearer %s", sender.Auth)
}

func (sender *Sender) Send(e interface{}) (map[string]any, error) {
	m := structs.Map(e)
	formValues := url.Values{}
	for k := range m {
		formValues.Set(k, cast.ToString(m[k]))
	}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, sender.url, nil)
	if err != nil {
		return nil, fmt.Errorf("new request err: %+v", err)
	}
	req.URL.RawQuery = formValues.Encode()
	req.Header.Set(httpx.ContentType, "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", sender.GetAuth())
	response, err := httpc.DoRequest(req)
	if err != nil {
		return nil, fmt.Errorf("request err: %+v", err)
	}
	defer response.Body.Close()

	// 解析 JSON 响应
	var data map[string]any
	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("decode response err: %+v", err)
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status err, status: %v, data: %v", response.StatusCode, data)
	}

	return data, nil
}
