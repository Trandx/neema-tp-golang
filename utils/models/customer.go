package models

type Customer struct {
	ID                   int     `json:"id" xorm:"'id' pk autoincr"`
	Customer_name        string  `xorm:"'customer_name'"`
	Street               string  `xorm:"'street'"`
	City                 string  `xorm:"'city'"`
	State                string  `xorm:"'state'"`
	Zip_code             string  `xorm:"'zip_code'"`
	Notes                string  `xorm:"'notes'"`
	Terms                uint    `xorm:"'terms'"`
	Account_number       string  `xorm:"'account_number'"`
	Tax_id               string  `xorm:"'tax_id'"`
	Balance              string  `xorm:"'balance'"`
	Credit_limit         string  `xorm:"'credit_limit'"`
	Is_active            bool    `xorm:"'is_active'"`
	Is_sub_agency        bool    `xorm:"'is_sub_agency'"`
	Opening_balance      string  `xorm:"'opening_balance'"`
	Language             string  `xorm:"'language'"`
	Slug                 int64   `xorm:"'slug'"`
	Id_currency          int64   `xorm:"'id_currency'"`
	Id_country           int64   `xorm:"'id_country'"`
	Irs_share_key        string  `xorm:"'irs_share_key'"`
	Currency_rate        float32 `xorm:"'currency_rate'"`
	Agency               string  `xorm:"'agency'"`
	Opening_balance_date string  `xorm:"opening_balance_date"`
	Avoid_deletion       bool    `xorm:"'avoid_deletion'"`
	Is_editable          bool    `xorm:"'is_editable'"`
	Alias                string  `xorm:"'alias'"`
	Already_used         int64   `xorm:"'already_used'"`
	Ab_key               string  `xorm:"'ab_key'"`
	Tmc_client_number    string  `xorm:"'tmc_client_number'"`
}

func (Customer) TableName() string {
	return "customer"
}
