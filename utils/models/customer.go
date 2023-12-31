package models

import (
	. "neema.co.za/rest/utils/models/_base"
)

type Customer struct {
	ID            int64  `json:"id" xorm:"'id' pk autoincr"`
	Customer_name string `json:"customer_name" xorm:"'customer_name'"`
	// Street            string  `xorm:"'street'"`
	// City              string  `xorm:"'city'"`
	State string `json:"state" xorm:"'state'"`
	// Zip_code          string  `xorm:"'zip_code'"`
	// Notes             string  `xorm:"'notes'"`
	// Terms             uint    `xorm:"'terms'"`
	Account_number string `json:"account_number" xorm:"'account_number'"`
	// Tax_id            string  `xorm:"'tax_id'"`
	// Balance           string  `xorm:"'balance'"`
	// Credit_limit      string  `xorm:"'credit_limit'"`
	// Is_active         bool    `xorm:"'is_active'"`
	// Is_sub_agency     bool    `xorm:"'is_sub_agency'"`
	// Opening_balance   string  `xorm:"'opening_balance'"`
	Slug        int64 `json:"slug" xorm:"'slug'"`
	Id_currency int64 `json:"id_currency" xorm:"'id_currency'"`
	Id_country  int64 `json:"id_country" xorm:"'id_country'"`
	// Irs_share_key     string  `xorm:"'irs_share_key'"`
	// Currency_rate     float32 `xorm:"'currency_rate'"`
	// Agency            string  `xorm:"'agency'"`
	// Avoid_deletion    bool    `xorm:"'avoid_deletion'"`
	// Is_editable       bool    `xorm:"'is_editable'"`
	Alias string `json:"alias" xorm:"'alias'"`
	// Already_used      int64   `xorm:"'already_used'"`
	Ab_key            string `json:"ab_key" xorm:"'ab_key'"`
	Tmc_client_number string `json:"tmc_client_number" xorm:"'tmc_client_number'"`
}

type CustomerPaginated struct {
	Data       *[]Customer `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

func (Customer) TableName() string {
	return "customer"
}
