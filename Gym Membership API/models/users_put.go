package models

import (
	"encoding/json"
	"time"
)

type Users2 struct {
	// ID          string      `json:"id"`
	// Tittle      string      `json:"tittle" binding:"required"`
	// Price       json.Number `json:"price" binding:"required"`
	// Description string      `json:"description" binding:"required"`
	// Rating      json.Number `json:"rating" binding:"required"`
	// Discount    json.Number `json:"discount" binding:"required"`
	// CreatedAt   time.Time
	// UpdateAt    time.Time

	ID        json.Number `json:"id"`
	Name      string      `json:"name"`
	Phone     string      `json:"phone"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time
}
