package models

import (
	"time"

	. "neema.co.za/rest/utils/models/_base"
)

type Invoice struct {
	ID             int64     `json:"id" xorm:"'id' pk autoincr"`
	Creation_date  time.Time `json:"creation_date" xorm:"'creation_date'"`
	Invoice_number string    `json:"invoice_number" xorm:"'invoice_number'"`
	Status         string    `json:"status" xorm:"'status'"`
	Due_date       string    `json:"due_date" xorm:"'due_date'"`
	Amount         string    `json:"amount" xorm:"'amount'"`
	Balance        string    `json:"balance" xorm:"'balance'"`
	// Purchase_order     string    `json:"purchase_order" xorm:"'purchase_order'"`
	// Customer_notes     string    `json:"customer_notes" xorm:"'customer_notes'"`
	Terms int8 `json:"terms" xorm:"'terms'"`
	// Terms_conditions   string    `json:"terms_conditions" xorm:"'terms_conditions'"`
	Credit_apply string `json:"credit_apply" xorm:"'credit_apply'"`
	// Rate               float64   `json:"rate" xorm:"'rate'"`
	Net_amount string `json:"net_amount" xorm:"'net_amount'"`
	// Tax_amount         string    `json:"tax_amount" xorm:"'tax_amount'"`
	Base_amount string `json:"base_amount" xorm:"'base_amount'"`
	// Slug               int64     `json:"slug" xorm:"'slug'"`
	Id_customer int64 `json:"id_customer" xorm:"'id_customer'"`
	// Credit_used        string    `json:"credit_used" xorm:"'credit_used'"`
	// Email              string    `json:"email" xorm:"'email'"`
	// Printed_name       string    `json:"printed_name" xorm:"'printed_name'"`
	// Hidden_field       string    `json:"hidden_field" xorm:"'hidden_field'"`
	// Hidden_identifier  string    `json:"hidden_identifier" xorm:"'hidden_identifier'"`
	// Already_used       int8      `json:"already_used" xorm:"'already_used'"`
	// Is_opening_balance bool      `json:"is_opening_balance" xorm:"'is_opening_balance'"`
	Tag string `json:"tag" xorm:"'tag'"`
	//Travel_Items []AirBooking `json:"air_booking_id" xorm:"has_many(id_customer)"`
}

type InvoicePaginated struct {
	Data       *[]Invoice `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type InvoiceWithRelation struct {
	Invoice
	AirBookings []AirBooking `json:"air_bookings"`
}

func (Invoice) TableName() string {
	return "invoice"
}
