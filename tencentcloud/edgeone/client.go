// Package edgeone provides the Tencent Cloud EdgeOne cache purge transport.
package edgeone

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/og-saas/framework/tencentcloud"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	teo "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/teo/v20220901"
	"github.com/zeromicro/go-zero/core/stringx"
)

type Result struct {
	JobID     string
	RequestID string
}

type Client struct{ sdk *teo.Client }

func NewClient(cfg tencentcloud.Config) (*Client, error) {
	if stringx.HasEmpty(cfg.SecretID) || stringx.HasEmpty(cfg.SecretKey) {
		return nil, fmt.Errorf("EdgeOne credential environment variables are missing")
	}
	p := profile.NewClientProfile()
	p.HttpProfile.Endpoint = "teo.tencentcloudapi.com"
	p.HttpProfile.ReqTimeout = cfg.ReqTimeout
	sdk, err := teo.NewClient(common.NewTokenCredential(cfg.SecretID, cfg.SecretKey, cfg.Token), "", p)
	if err != nil {
		return nil, err
	}
	return &Client{sdk: sdk}, nil
}

func (c *Client) PurgeDirectories(ctx context.Context, zone string, targets []string) (Result, error) {
	if strings.TrimSpace(zone) == "" || len(targets) == 0 {
		return Result{}, fmt.Errorf("zone and targets are required")
	}
	req := teo.NewCreatePurgeTaskRequest()
	req.ZoneId = common.StringPtr(zone)
	req.Type = common.StringPtr("purge_prefix")
	req.Method = common.StringPtr("delete")
	req.Targets = common.StringPtrs(targets)
	resp, err := c.sdk.CreatePurgeTaskWithContext(ctx, req)
	if err != nil {
		return Result{}, err
	}
	if resp == nil || resp.Response == nil {
		return Result{}, fmt.Errorf("empty EdgeOne response")
	}
	r := resp.Response
	result := Result{JobID: stringValue(r.JobId), RequestID: stringValue(r.RequestId)}
	if len(r.FailedList) != 0 {
		details, _ := json.Marshal(r.FailedList)
		return result, fmt.Errorf("EdgeOne partial failure: request_id=%s failures=%s", result.RequestID, details)
	}
	if result.JobID == "" {
		return result, fmt.Errorf("EdgeOne returned no job ID: request_id=%s", result.RequestID)
	}
	return result, nil
}

func stringValue(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
