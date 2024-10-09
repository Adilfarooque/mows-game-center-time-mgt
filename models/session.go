package models

import "time"

// type Session struct {
// 	ID         int       `json:"id" gorm:"primaryKey"`
// 	GameID     int       `json:"game_id" gorm:"not null"`
// 	UserID     int       `json:"user_id" gorm:"not null"`
// 	StartTime  time.Time `json:"start_time" gorm:"not null"`
// 	EndTime    time.Time `json:"end_time"`
// 	Status     string    `json:"status"`      // e.g., "ongoing", "completed", "canceled"
// 	TotalPrice float64   `json:"total_price"` // Total cost of the session
// 	CreatedAt  time.Time `json:"created_at"`
// 	UpdatedAt  time.Time `json:"updated_at"`
// }

type Session struct {
	ID        string        `gorm:"primaryKey"`
	GameID    string        `json:"game_id"`
	UserID    string        `json:"user_id"`
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
	Status    string        `json:"status"`
	Duration  time.Duration `json:"duration"`
}

type SessionRequest struct {
	GameID    string    `json:"game_id" binding:"required"`
	UserID    string    `json:"user_id" binding:"required"`
	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
}

type SessionRescheduleRequest struct {
	GameID       int       `json:"game_id" binding:"required"`
	NewStartTime time.Time `json:"new_start_time" binding:"required"`
	NewEndTime   time.Time `json:"new_end_time" binding:"required"`
}
