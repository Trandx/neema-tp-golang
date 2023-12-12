package repository

import (
	"math"

	. "neema.co.za/rest/utils/logger"
	. "neema.co.za/rest/utils/models"
	. "neema.co.za/rest/utils/models/_base"
)

func (r *Repository) DeleteInvoice(invoice *Invoice) (string, error) {

	_, err := r.Delete(invoice)

	if err != nil {
		Logger.Error(err.Error())
		return "", err
	}

	Logger.Info("Invoice deleted")

	return "Invoice deleted successfully", nil
}

func (r *Repository) CreateInvoice(invoice *Invoice) (*Invoice, error) {

	_, err := r.Insert(invoice)

	if err != nil {
		Logger.Error(err.Error())
		return nil, err
	}

	Logger.Info("Invoice saved")

	return invoice, nil
}

func (r *Repository) GetInvoices(params *Params) (*Invoice, *InvoicePaginated, error) {
	var invoice []Invoice

	if params.Id > 0 {

		err := r.Find(&invoice, &Invoice{ID: params.Id})

		if err != nil {
			Logger.Error(err.Error())
			return nil, nil, err
		}

		if len(invoice) == 0 {
			return nil, nil, nil
		}

		return &invoice[0], nil, nil
	}

	offset := (params.Page - 1) * params.PageSize
	totalItem, err := r.Count(&Invoice{})
	totalPage := math.Ceil(float64(totalItem) / float64(params.PageSize))

	if err != nil {
		return nil, nil, err
	}

	err = r.Limit(params.PageSize, offset).Find(&invoice)
	if err != nil {
		return nil, nil, err
	}
	var invoicePaginated InvoicePaginated

	invoicePaginated.Data = &invoice
	invoicePaginated.Pagination.CurrentPage = int64(params.Page)
	invoicePaginated.Pagination.Size = int64(params.PageSize)
	invoicePaginated.Pagination.TotalPages = int64(totalPage)
	invoicePaginated.Pagination.TotalItem = totalItem

	return nil, &invoicePaginated, nil

}
