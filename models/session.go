package models

type Session struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	GameID     int     `json:"game_id"`
	Start      string  `json:"start"`
	End        string  `json:"end"`
	Status     string  `json:"status"`      // e.g., "ongoing", "completed", "canceled"
	TotalPrice float64 `json:"total_price"` // price for the session based on game duration
}
