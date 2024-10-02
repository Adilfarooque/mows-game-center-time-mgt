package models

import "time"

type Game struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title" gorm:"not null"`
	Category   string    `json:"category"`   // e.g., "racing", "arcade", "shooting"
	Rating     float64   `json:"rating"`     // Average rating
	Popularity int       `json:"popularity"` // Play count
	Price      float64   `json:"price"`
	ImageURL   string    `json:"image_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
