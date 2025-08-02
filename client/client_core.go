package client

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"resty.dev/v3"
)

func (c *client) getHttpRequest(ctx context.Context) *resty.Request {
	return httpClient.SetBaseURL(ApiUrl).
		R().
		SetForceResponseContentType("application/json").
		SetDebug(c.isDebug).
		SetContext(ctx)
}

func (c *client) getAccessToken() (string, error) {
	if c.authToken == nil {
		return "", errors.New("auth token not set")
	}

	if c.authToken.IsExpired() {
		return "", errors.New("auth token expired")
	}

	return c.authToken.AccessToken(), nil
}

func (c *client) get(ctx context.Context, apiName ApiURL, params any, result any) error {
	var ak string

	if apiName.RequiresToken() || c.forceWithToken {
		var err error

		ak, err = c.getAccessToken()
		if err != nil {
			return err
		}
	}

	values, err := query.Values(params)
	if err != nil {
		return err
	}

	var (
		respErr = new(RespErr)
	)

	if _, err = c.getHttpRequest(ctx).
		SetResult(result).
		SetError(respErr).
		WithContext(ctx).
		SetHeaders(signatureHeader(values, ak)).
		Get(fmt.Sprintf("%s?%s", apiName, values.Encode())); err != nil {
		return err
	}

	if respErr.HasError() {
		return respErr
	}

	return nil
}

func (c *client) WithDebug(flags ...bool) Client {
	if len(flags) > 0 {
		c.isDebug = flags[0]
	} else {
		c.isDebug = true
	}

	return c
}

func (c *client) WithToken(token AuthToken) Client {
	c.authToken = token

	return c
}

func (c *client) WithForceWithToken(flags ...bool) Client {
	if len(flags) > 0 {
		c.forceWithToken = flags[0]
	} else {
		c.forceWithToken = true
	}

	return c
}

func (c *client) WithClient(client *resty.Client) Client {
	c.httpClient = client

	return c
}
