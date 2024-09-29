package models

type Payment struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	Amount    float64 `json:"amount"`
	Method    string  `json:"method"` //  e.g., "credit card", "cash", "wallet"
	Timestamp string  `json:"timestamp"`
	SessionID int     `json:"session_id"` // reference to the session
}
