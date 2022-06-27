package entity

import "time"

type Class_Booking_Details struct {
	ID                 int
	Class_Booking_ID   int
	Class_ID           int
	Class_Name         string
	Class_Description  string
	Class_Meet_Link    string
	Class_Category     int
	Class_Image        string
	Class_Price        float64
	Class_Date         time.Time
	Class_Start        time.Time
	Class_End          time.Time
	Class_Trainer_Name string
	CreatedAt          time.Time
	UpdateAt           time.Time
	DeleteAt           time.Time
}
