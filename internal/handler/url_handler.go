package handler

import (
	"net/http"

	"7charUrl/internal/service"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(s *service.URLService) *URLHandler {
	return &URLHandler{service: s}
}

func (h *URLHandler) ShortenURL(c *gin.Context) {
	var body struct {
		URL string `json:"url"`
	}

	if err := c.BindJSON(&body); err != nil || body.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	short := h.service.Shorten(body.URL)

	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:8080/" + short,
	})
}

func (h *URLHandler) Redirect(c *gin.Context) {
	code := c.Param("code")

	original, ok := h.service.Resolve(code)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "short URL not found"})
		return
	}

	c.Redirect(http.StatusFound, original)
}
