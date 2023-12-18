package models

import (
	. "neema.co.za/rest/utils/models/_base"
)

type Payment struct {
	ID int64 `json:"id" xorm:"'id' pk autoincr"`
}

type PaymentPaginated struct {
	Data       *[]PaymentWithRelation `json:"data"`
	Pagination Pagination             `json:"pagination"`
}

type PaymentWithRelation struct {
	Payment
	AirBookings []AirBooking `json:"air_bookings" xorm:"'air_bookings'"`
	Customer    *Customer    `json:"customer" xorm:"'customer'"`
}

func (Payment) TableName() string {
	return "payment"
}
