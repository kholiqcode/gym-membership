package entity

import "time"

type Member_Types struct {
	ID          int
	Name        string
	Description string
	Image       string
	Duration    int
	Price       float64
	CreatedAt   time.Time
	UpdateAt    time.Time
	DeleteAt    time.Time
}
