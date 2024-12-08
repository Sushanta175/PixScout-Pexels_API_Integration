package client

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/Sushanta175/Go_Pexels_API/config"
	"github.com/Sushanta175/Go_Pexels_API/models"
)

func (c *Client) SearchPhotos(query string, perPage, page int) (*models.PhotoSearchResult, error) {
	url := fmt.Sprintf(config.PhotoApi+"search?query=%s&per_page=%s&page=%s", query, perPage, page)

	resp, err := c.RequestDoWithAuth("GET", url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	defer func() {
		if resp.Body != nil {
			defer resp.Body.Close()
		}
	}()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("API error: received status code %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %v", err)
	}

	var result models.PhotoSearchResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("Error Unmarshalling JSON: %v", err)
	}

	return &result, nil
}
