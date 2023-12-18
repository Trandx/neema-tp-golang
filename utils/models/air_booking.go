package models

import (
	. "neema.co.za/rest/utils/models/_base"
)

type AirBooking struct {
	ID               int64   `json:"id" xorm:"'id' pk autoincr"`
	Product_type     string  `json:"product_type" xorm:"'product_type'"`
	Transaction_type string  `json:"transaction_type" xorm:"'transaction_type'"`
	Total_price      float64 `json:"total_price" xorm:"'total_price'"`
	Status           string  `json:"status" xorm:"'status'"`
	Id_invoice       int64   `json:"id_invoice" xorm:"'id_invoice'"`
	Traveler_name    string  `json:"traveler_name" xorm:"'traveler_name'"`
	Issuing_date     string  `json:"issuing_date" xorm:"'issuing_date'"`
	Id_customer      int64   `json:"id_customer" xorm:"'id_customer'"`
	Fop              string  `json:"fop" xorm:"'fop'"`
	Ticket_number    int64   `json:"ticket_number" xorm:"'ticket_number'"`
	Id_airline       int64   `json:"id_airline" xorm:"'id_airline'"`
	Id_product       int64   `json:"id_product" xorm:"'id_product'"`
	Id_agent_sign    int64   `json:"id_agent_sign" xorm:"'id_agent_sign'"`
	Coz_of_supplier  bool    `json:"coz_of_supplier" xorm:"'coz_of_supplier'"`
}

type AirBookingPaginated struct {
	Data       *[]AirBooking `json:"data"`
	Pagination Pagination    `json:"pagination"`
}

func (AirBooking) TableName() string {
	return "air_booking"
}
