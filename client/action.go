// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "goa simple sample": action Resource Client
//
// Command:
// $ goagen
// --design=github.com/testGoa/design
// --out=$(GOPATH)/src/github.com/testGoa
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// PingActionPath computes a request path to the ping action of action.
func PingActionPath() string {

	return fmt.Sprintf("/api/v1/action/ping")
}

// サーバーとの導通確認
func (c *Client) PingAction(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewPingActionRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPingActionRequest create the request corresponding to the ping action endpoint of the action resource.
func (c *Client) NewPingActionRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
