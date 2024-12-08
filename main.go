package main

import (
	"fmt"
	"log"

	"github.com/Sushanta175/Go_Pexels_API/client"
	"github.com/Sushanta175/Go_Pexels_API/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	cfg := config.LoadConfig()

	var c = client.NewClient(cfg.ApiToken)

	result, err := c.SearchPhotos("waves", 15, 1)
	if err != nil {
		log.Printf("Search Error %v", err)
	}
	if result.Page == 0 {
		log.Printf("Something wrong in the search")
	}
	fmt.Println(result)
}
