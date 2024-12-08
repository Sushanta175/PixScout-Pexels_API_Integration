package client

import (
	"encoding/json"
	"fmt"
	"io"
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

func (c *Client) performRequestWithAuth(method, url string) (*http.Response, error) {
	resp, err := c.RequestDoWithAuth(method, url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		defer resp.Body.Close()
		return nil, fmt.Errorf("api error: received status code %d", resp.StatusCode)
	}

	return resp, nil
}

func parseResponseBody[T any](resp *http.Response, result *T) error {
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		return fmt.Errorf("error Unmarshalling JSON: %v", err)
	}

	return nil
}

func (c *Client) RemainingRequests() int32 {
	return c.RemainingTimes
}
