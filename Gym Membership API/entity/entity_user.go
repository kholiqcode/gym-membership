package entity

import "time"

type User struct {
	ID        int
	Name      string
	Phone     string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time
}
