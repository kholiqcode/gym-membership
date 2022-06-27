package models

import (
	"encoding/json"
	"time"
)

type Users struct {

	ID        json.Number `json:"id"`
	Name      string      `json:"name" binding:"required"`
	Phone     string      `json:"phone" binding:"required"`
	Email     string      `json:"email" binding:"required"`
	Password  string      `json:"password" binding:"required"`
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time
}
