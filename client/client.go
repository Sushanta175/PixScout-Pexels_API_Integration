package client

import (
	"fmt"
	"net/http"
	"strconv"
)

type Client struct {
	Token          string
	hc             http.Client
	RemainingTimes int32
}

func NewClient(token string) *Client {
	c := http.Client{}
	return &Client{Token: token, hc: c}
}

func (c *Client) RequestDoWithAuth(method, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %w", err)
	}

	req.Header.Add("Authorization", c.Token)

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if rateLimit := resp.Header.Get("X-Ratelimit-Remaining"); rateLimit != "" {
		if times, err := strconv.Atoi(rateLimit); err == nil {
			c.RemainingTimes = int32(times)
		} else {
			fmt.Printf("warning: Invalid X-Ratelimit-Remaining Header: %v", rateLimit)
		}
	}

	return resp, nil
}
