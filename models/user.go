package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`    //"player" or "admin"
	Credits  int    `json:"credits"` //represent the balance that users use to book game sessions.
}

