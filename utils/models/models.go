package models

type Games struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Duration string `json:"duration"` // in minites
	Category string `json:"category"` //game category (arcade,racing)
	Rating float64 `json:"rating"` //user rating for the game
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Role string `json:"role"` //"player" or "admin"
	Credits int `json:"credits"` //represent the balance that users use to book game sessions.
}

type Session struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	GameID int    `json:"game_id"`
	Start  string `json:"start"`
	End    string `json:"end"`
	Status string `json:"status"` // e.g., "ongoing", "completed", "canceled"
	TotalPrice float64 `json:"total_price"` // price for the session based on game duration
}

type Payment struct{
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Amount float64 `json:"amount"`
	Method string `json:"method"` //  e.g., "credit card", "cash", "wallet"
	Timestamp string	`json:"timestamp"`
	SessionID int `json:"session_id"` // reference to the session

}

