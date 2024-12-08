package client

import (
	"fmt"
	"math/rand"

	"github.com/Sushanta175/Go_Pexels_API/config"
	"github.com/Sushanta175/Go_Pexels_API/models"
)

func (c *Client) SearchVideo(query string, perPage, page int) (*models.VideoSearchResult, error) {
	url := fmt.Sprintf(config.VideoApi+"search?query=%s&per_page=%d&page=%d", query, perPage, page)

	resp, err := c.performRequestWithAuth("GET", url)
	if err != nil {
		return nil, err
	}

	var result models.VideoSearchResult
	err = parseResponseBody(resp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PopularVideo(perPage, page int) (*models.PopularVideos, error) {
	url := fmt.Sprintf(config.VideoApi+"popular?per_page=%d&page=%d", perPage, page)

	resp, err := c.performRequestWithAuth("GET", url)
	if err != nil {
		return nil, err
	}

	var result models.PopularVideos
	err = parseResponseBody(resp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetRandomVideo() (*models.Video, error) {
	page := rand.Intn(1001)
	result, err := c.PopularVideo(1, page)

	if err == nil && len(result.Videos) == 1 {
		return &result.Videos[0], nil
	}

	return nil, err
}
