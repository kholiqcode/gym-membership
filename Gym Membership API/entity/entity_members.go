package entity

import "time"

type Members struct {
	ID             int
	Member_Type_ID int
	Name           string
	Phone          string
	Email          string
	Password       string
	Is_Active      string
	CreatedAt      time.Time
	UpdateAt       time.Time
	DeleteAt       time.Time
}
