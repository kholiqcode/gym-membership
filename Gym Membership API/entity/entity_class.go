package entity

import "time"

type Class struct {
	ID          int
	Trainer_ID  int
	Name        string
	Description string
	Category    string
	Image       string
	Meet_Link   string
	Price       float64
	Date        time.Time
	Start_Time  time.Time
	End_Time    time.Time
	CreatedAt   time.Time
	UpdateAt    time.Time
	DeleteAt    time.Time
}
