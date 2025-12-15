package models

import "time"

type URL struct {
	ID          uint      `gorm:"primaryKey"`
	ShortURL    string    `gorm:"size:10;uniqueIndex;not null"`
	LongURL     string    `gorm:"type:text;not null"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	ExpiryTime  time.Time
	Count       int64 `gorm:"default:0"`
}
