package models

type Games struct {
	ID       int     `json:"id" gorm:"primarykey"`
	Name     string  `json:"name"`
	Duration string  `json:"duration"` // in minites
	Category string  `json:"category"` //game category (arcade,racing)
	Rating   float64 `json:"rating"`   //user rating for the game
}
