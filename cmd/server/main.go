package main

import (
	"7charUrl/internal/config"
	"7charUrl/internal/handler"
	"7charUrl/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	r := gin.Default()

	urlService := service.NewURLService()
	urlHandler := handler.NewURLHandler(urlService)

	r.POST("/shorten", urlHandler.ShortenURL)
	r.POST("/custom", urlHandler.CustomURL)
	r.GET("/:code", urlHandler.Redirect)

	r.Run(":8080")
}
