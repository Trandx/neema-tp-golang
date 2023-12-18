package repository

import (
	"math"
	"strconv"
	"strings"
	"time"

	. "neema.co.za/rest/utils/logger"
	. "neema.co.za/rest/utils/models"
	. "neema.co.za/rest/utils/models/_base"
)

// func getStructKeys() []string {

// 	test := AirBooking{}
// 	e := reflect.ValueOf(&test).Elem()

// 	var structKeys []string

// 	for i := 0; i < e.NumField(); i++ {
// 		keyName := e.Type().Field(i).Name
// 		structKeys = append(structKeys, keyName)
// 		// varType := e.Type().Field(i).Type
// 		// varValue := e.Field(i).Interface()
// 		//fmt.Printf("%v %v %v\n", varName,varType,varValue)
// 	}

// 	return structKeys
// }

func (r *Repository) DeleteInvoice(invoice *Invoice) (string, error) {

	_, err := r.Delete(invoice)

	if err != nil {
		Logger.Error(err.Error())
		return "", err
	}

	Logger.Info("Invoice deleted")

	return "Invoice deleted successfully", nil
}

func (r *Repository) CreateInvoice(invoiceWithRelation *InvoiceWithRelation) (*InvoiceWithRelation, error) {

	// make a transaction
	session := r.NewSession()

	//defer r.Close()

	invoice := invoiceWithRelation.Invoice

	_airBookings := invoiceWithRelation.AirBookings

	invoice.Creation_date = time.Now()
	invoice.Credit_apply = 0
	invoice.Balance = invoice.Amount
	invoice.Base_amount = invoice.Amount
	invoice.Tag = "1"
	invoice.Status = "unpaid"

	// add Begin() before any action
	err := session.Begin()

	if err != nil {
		Logger.Error(err.Error())
		return nil, err
	}

	// operation start
	_, err = session.Insert(&invoice)

	if err != nil {
		session.Rollback()
		Logger.Error(err.Error())
		return nil, err
	}

	var IdList []string
	// update id of customer into each airbooking items
	for _, airbooairBooking := range _airBookings {
		IdList = append(IdList, strconv.FormatInt(airbooairBooking.ID, 10))
	}

	//structKeys := getStructKeys()
	//columns := strings.ToLower(strings.Join(structKeys, ","))
	ids := strings.Join(IdList, ",")

	query := "UPDATE air_booking SET status='invoiced', id_invoice='"
	query += strconv.FormatInt(invoice.ID, 10)
	query += "' "
	query += "WHERE id IN (" + ids + ") RETURNING *"

	var airBookings []AirBooking

	err = session.SQL(query).Find(&airBookings)

	//log.Println(&airBookings)

	//session.Rollback() // to unsave transaction

	if err != nil {
		session.Rollback()
		Logger.Error(err.Error())
		return nil, err
	}

	Logger.Info("Invoice saved")

	//update InvoiceWithRelation model datas
	invoiceWithRelation.Invoice = invoice
	invoiceWithRelation.AirBookings = airBookings

	// add Commit() after all actions
	err = session.Commit()
	if err != nil {
		Logger.Error(err.Error())
		return nil, err
	}

	return invoiceWithRelation, nil
}

func (r *Repository) GetInvoices(params *Params) (*InvoiceWithRelation, *InvoicePaginated, error) {
	var invoiceWithRelation []InvoiceWithRelation

	if params.Id > 0 {

		err := r.Join("INNER", "air_booking", "air_booking.id_invoice = invoice.id").Join("INNER", "customer", "customer.id = invoice.id_customer").Find(&invoiceWithRelation, &Invoice{ID: params.Id})

		if err != nil {
			Logger.Error(err.Error())
			return nil, nil, err
		}

		if len(invoiceWithRelation) == 0 {
			return nil, nil, nil
		}

		return &invoiceWithRelation[0], nil, nil
	}

	offset := (params.Page - 1) * params.PageSize
	totalItem, err := r.Count(&Invoice{})
	totalPage := math.Ceil(float64(totalItem) / float64(params.PageSize))

	if err != nil {
		return nil, nil, err
	}

	query := "SELECT * FROM invoice JOIN air_booking ON air_booking.id_invoice = invoice.id;"

	err = r.SQL(query).Limit(params.PageSize, offset).Find(&invoiceWithRelation)

	// err = r.Limit(params.PageSize, offset).Join("ON ", "air_booking", "air_booking.id_invoice = invoice.id").Find(&invoiceWithRelation)
	if err != nil {
		return nil, nil, err
	}
	var invoicePaginated InvoicePaginated

	invoicePaginated.Data = &invoiceWithRelation
	invoicePaginated.Pagination.CurrentPage = int64(params.Page)
	invoicePaginated.Pagination.Size = int64(params.PageSize)
	invoicePaginated.Pagination.TotalPages = int64(totalPage)
	invoicePaginated.Pagination.TotalItem = totalItem

	return nil, &invoicePaginated, nil

}
