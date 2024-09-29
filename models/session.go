package models

import "time"

type Session struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	GameID     int       `json:"game_id" gorm:"not null"`
	UserID     int       `json:"user_id" gorm:"not null"`
	StartTime  time.Time `json:"start_time" gorm:"not null"`
	EndTime    time.Time `json:"end_time"`
	Status     string    `json:"status"`      // e.g., "ongoing", "completed", "canceled"
	TotalPrice float64   `json:"total_price"` // Total cost of the session
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
