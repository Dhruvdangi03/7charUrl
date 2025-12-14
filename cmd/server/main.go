package main

import (
	"7charUrl/internal/handler"
	"7charUrl/internal/service"
	"7charUrl/internal/store"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	memoryStore := store.NewMemoryStore()
	urlService := service.NewURLService(memoryStore)
	urlHandler := handler.NewURLHandler(urlService)

	r.POST("/shorten", urlHandler.ShortenURL)
	r.GET("/:code", urlHandler.Redirect)

	r.Run(":8080")
}
