package main

import (
	"fmt"

	"github.com/Sushanta175/Go_Pexels_API/client"
	"github.com/Sushanta175/Go_Pexels_API/config"
)

func main() {
	cfg := config.LoadConfig()

	var c = client.NewClient(cfg.ApiToken)

	result, err := c.SearchPhotos("waves", 15, 1)
	if err != nil {
		fmt.Errorf("Search Error %v", err)
	}
	if result.Page == 0 {
		fmt.Errorf("Something wrong in the search")
	}
	fmt.Println(result)
}
