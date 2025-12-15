package handler

import (
	"net/http"

	"7charUrl/internal/service"

	"github.com/gin-gonic/gin"
)

var HOSTURL string

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

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	HOSTURL = scheme + "://" + c.Request.Host

	if err := c.BindJSON(&body); err != nil || body.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	short := h.service.Shorten(body.URL)

	c.JSON(http.StatusOK, gin.H{
		"short_url": HOSTURL + "/" + short,
	})
}

func (h *URLHandler) CustomURL(c *gin.Context) {
	var body struct {
		URL    string `json:"url"`
		Custom string `json:"custom"`
	}

	if err := c.BindJSON(&body); err != nil || body.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	custom, ok := h.service.Custom(body.URL, body.Custom)

	if !ok {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"Error": custom,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short_url": HOSTURL + "/" + custom,
	})
}

func (h *URLHandler) Redirect(c *gin.Context) {
	code := c.Param("code")

	original, ok := h.service.Resolve(code)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "short URL not found"})
		return
	}

	c.Redirect(http.StatusFound, "https://"+original)
}
