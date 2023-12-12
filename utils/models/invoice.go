package models

import (
	"time"

	. "neema.co.za/rest/utils/models/_base"
)

type Invoice struct {
	ID                 int64 `json:"id" xorm:"'id' pk autoincr"`
	Creation_date      time.Time
	Invoice_number     string
	Status             string
	Due_date           string
	Amount             string
	Balance            string
	Purchase_order     string
	Customer_notes     string
	Terms              int8
	Terms_conditions   string
	Credit_apply       string
	Rate               float64
	Net_amount         string
	Tax_amount         string
	Base_amount        string
	Slug               int64
	Id_customer        int64
	Credit_used        string
	Email              string
	Printed_name       string
	Hidden_field       string
	Hidden_identifier  string
	Already_used       int8
	Is_opening_balance bool
	Tag                string `xorm:"tag"`
	//AirBookingIds      []int64 `xorm:"'air_booking_ids'"`
}

type InvoicePaginated struct {
	Data       *[]Invoice `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func (Invoice) TableName() string {
	return "invoice"
}
