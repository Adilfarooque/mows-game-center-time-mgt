package models

import "time"

type Admin struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	UserID      int       `json:"user_id" gorm:"not null"` // Link to User model
	Permissions []string  `json:"permissions"`             // e.g., "manage_games", "manage_users"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DashBoardUser struct {
	TotalUsers int `json:"Totaluser"`
	BlockUser  int `json:"Blockuser"`
}

type DashBoardGames struct {
	TotalGames        int `json:"Totalusers"`
	OutofStockProduct int `json:"Outofstock"`
}

