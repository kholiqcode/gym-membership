package entity

import "time"

type Trainers struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Avatar    string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time
}
