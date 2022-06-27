package entity

import "time"

type Merber_Orders struct {
	ID                   int
	Member_ID            int
	Member_Type_ID       int
	Invoice_No           string
	Member_Name          string
	Member_Email         string
	Member_Phone         string
	Member_Type_Name     string
	Member_Type_Image    string
	Member_Type_Duration int
	Member_Type_Price    float64
	Payment_Method       string
	Status               int
	Total                float64
	CreatedAt            time.Time
	UpdateAt             time.Time
	DeleteAt             time.Time
}
