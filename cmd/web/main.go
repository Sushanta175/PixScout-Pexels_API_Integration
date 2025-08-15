package main

import (
	"net/http"
	"os"

	"github.com/Sushanta175/Go_Pexels_API/client"
	"github.com/Sushanta175/Go_Pexels_API/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	cfg := config.LoadConfig()
	c := client.NewClient(cfg.ApiToken)

	r := gin.Default()
	r.LoadHTMLGlob("web/templates/*")
	r.Static("/static", "./web/static")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/api/photos", func(ctx *gin.Context) {
		q := ctx.Query("q")
		if q == "" {
			q = "nature"
		}
		res, err := c.SearchPhotos(q, 30, 1)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, res)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
