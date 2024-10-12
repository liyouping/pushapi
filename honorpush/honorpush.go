package honorpush

import (
	"context"
	"fmt"
	"github.com/liyouping/pushapi/httputil"
	"net/http"
	"time"
)

type Client struct {
	httpClient        *http.Client
	host              string
	appId             string //应用ID, 开发者平台开通该应用PUSH服务后应用的APP ID
	clientId          string
	clientSecret      string
	authToken         string
	authTokenExpireAt int64
}

func NewClient(appId, clientId, clientSecret string) *Client {
	return &Client{
		host:         Host,
		appId:        appId,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (c *Client) SetHost(host string) {
	c.host = host
}

func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

func (c *Client) auth(ctx context.Context) (string, error) {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	if c.authToken != "" && c.authTokenExpireAt > now {
		return c.authToken, nil
	}

	req := &AuthReq{
		GrantType:    "client_credentials", //固定值为:client_credentials
		ClientID:     c.clientId,
		ClientSecret: c.clientSecret,
	}
	params := httputil.StructToUrlValues(req)
	res := &AuthRes{}

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	code, resBody, err := httputil.PostForm(ctx, c.httpClient, c.host+AuthURL, params, res, headers)
	if err != nil {
		return "", fmt.Errorf("code=%d body=%s err=%v", code, resBody, err)
	}

	if code != http.StatusOK || res.AccessToken == "" {
		return "", fmt.Errorf("code=%d body=%s", code, resBody)
	}

	c.authToken = res.AccessToken
	c.authTokenExpireAt = now + res.ExpiresIn - 60 //提前60秒刷新
	return c.authToken, nil
}

func (c *Client) Send(req *SendReq) (*SendRes, error) {
	return c.SendWithContext(context.Background(), req)
}

func (c *Client) SendWithContext(ctx context.Context, req *SendReq) (*SendRes, error) {
	res := &SendRes{}

	token, err := c.auth(ctx)
	if err != nil {
		return nil, err
	}
	pushUrl := fmt.Sprintf("%s/api/v1/%s/sendMessage", PushHost, c.appId)
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
		"timestamp":     fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond)),
	}
	code, resBody, err := httputil.PostJSON(ctx, c.httpClient, pushUrl, req, res, headers)
	if err != nil {
		return nil, fmt.Errorf("code=%d body=%s err=%v", code, resBody, err)
	}

	if code != http.StatusOK || res.Code != 200 {
		return nil, fmt.Errorf("code=%d body=%s", code, resBody)
	}

	return res, nil
}
