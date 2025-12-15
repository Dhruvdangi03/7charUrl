package store

import (
	"7charUrl/internal/config"
	"7charUrl/internal/models"
	"log"
	"time"
)

func CreateURL(shortURL string, longURL string) error {
	url := models.URL{
		ShortURL:    shortURL,
		LongURL:     longURL,
		CreatedTime: time.Now(),
		ExpiryTime:  time.Now().AddDate(0, 1, 0),
		Count:       0,
	}
	return config.DB.Create(&url).Error
}

func GetURL(shortURL string) (*models.URL, error) {
	var url models.URL
	result := config.DB.Where("short_url = ?", shortURL).First(&url)
	err := UpdateURLCount(url.ShortURL, url.Count)

	if err != nil {
		log.Fatal("❌ There is no entry in database:", err)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &url, nil
}

func UpdateURL(shortURL string, newLongURL string) error {
	return config.DB.Model(&models.URL{}).
		Where("short_url = ?", shortURL).
		Update("long_url", newLongURL).
		Error
}

func UpdateURLCount(shortURL string, count int64) error {
	return config.DB.Model(&models.URL{}).
		Where("short_url = ?", shortURL).
		Update("count", count+1).
		Error
}

func DeleteURL(shortURL string) error {
	return config.DB.Where("short_url = ?", shortURL).Delete(&models.URL{}).Error
}
