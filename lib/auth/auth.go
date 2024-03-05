package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	cl   *http.Client
	addr string
}

func New(addr string) *Client {
	return &Client{
		cl:   &http.Client{},
		addr: addr,
	}
}

func (c *Client) validate(ctx context.Context, jwtToken string) (*UserInfo, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.addr+"/validate?token="+jwtToken, nil)
	if err != nil {
		return nil, fmt.Errorf("can not create auth request: %w", err)
	}

	resp, err := c.cl.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can not execute auth request: %w", err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can not read auth responce: %w", err)
	}
	defer resp.Body.Close()

	var userInfo UserInfo

	err = json.Unmarshal(data, &userInfo)
	if err != nil {
		return nil, fmt.Errorf("can not unmarshal auth responce: %w", err)
	}

	return &userInfo, nil
}
