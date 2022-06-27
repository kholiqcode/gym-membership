package entity

import "time"

type Class_Bookings struct {
	ID             int
	Members_ID     int
	Invoice_No     string
	Member_Name    string
	Member_Email   string
	Member_Phone   string
	Payment_Method string
	Note           string
	Status         int
	Total          float64
	CreatedAt      time.Time
	UpdateAt       time.Time
	DeleteAt       time.Time
}
