package models

type Games struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Duration string `json:"duration"` // in minites
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Session struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	GameID int    `json:"game_id"`
	Start  string `json:"start"`
	End    string `json:"end"`
}
