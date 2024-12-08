package client

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/Sushanta175/Go_Pexels_API/config"
	"github.com/Sushanta175/Go_Pexels_API/models"
)

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

func (c *Client) SearchPhotos(query string, perPage, page int) (*models.PhotoSearchResult, error) {
	url := fmt.Sprintf(config.PhotoApi+"search?query=%s&per_page=%d&page=%d", query, perPage, page)

	resp, err := c.performRequestWithAuth("GET", url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	var result models.PhotoSearchResult
	err = parseResponseBody(resp, &result)
	return &result, err
}

func (c *Client) CuratedPhotos(per_page, page int) (*models.CuratedPhotosResult, error) {
	url := fmt.Sprintf(config.PhotoApi+"curated?per_page=%d&page=%d", per_page, page)

	resp, err := c.performRequestWithAuth("GET", url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	var result models.CuratedPhotosResult
	err = parseResponseBody(resp, &result)
	return &result, err
}

func (c *Client) GetPhoto(id int32) (*models.Photo, error) {
	url := fmt.Sprintf(config.PhotoApi+"photos/%d", id)

	resp, err := c.performRequestWithAuth("GET", url)
	if resp != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	var result models.Photo
	err = parseResponseBody(resp, &result)
	return &result, err
}

func (c *Client) GetRandomPhoto() (*models.Photo, error) {
	page := rand.Intn(1001)

	result, err := c.CuratedPhotos(1, page)
	if err == nil && len(result.Photos) == 1 {
		return &result.Photos[0], nil
	}

	return nil, err
}
