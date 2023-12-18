package repository

import (
	"math"

	. "neema.co.za/rest/utils/logger"
	. "neema.co.za/rest/utils/models"
	. "neema.co.za/rest/utils/models/_base"
)

// func (r *Repository) CreatePayment(payment *Payment) (*PaymentWithRelation, error)  {
//   //return &payment, err
// }

func (r *Repository) GetPayments(params *Params) (*PaymentWithRelation, *PaymentPaginated, error) {
	var paymentWithRelation []PaymentWithRelation

	if params.Id > 0 {

		err := r.Join("INNER", "air_booking", "air_booking.id_Payment = Payment.id").Join("INNER", "customer", "customer.id = Payment.id_customer").Find(&paymentWithRelation, &Payment{ID: params.Id})

		if err != nil {
			Logger.Error(err.Error())
			return nil, nil, err
		}

		if len(paymentWithRelation) == 0 {
			return nil, nil, nil
		}

		return &paymentWithRelation[0], nil, nil
	}

	offset := (params.Page - 1) * params.PageSize
	totalItem, err := r.Count(&Payment{})
	totalPage := math.Ceil(float64(totalItem) / float64(params.PageSize))

	if err != nil {
		return nil, nil, err
	}

	query := "SELECT * FROM Payment JOIN air_booking ON air_booking.id_Payment = Payment.id;"

	err = r.SQL(query).Limit(params.PageSize, offset).Find(&paymentWithRelation)

	// err = r.Limit(params.PageSize, offset).Join("ON ", "air_booking", "air_booking.id_Payment = Payment.id").Find(&PaymentWithRelation)
	if err != nil {
		return nil, nil, err
	}
	var PaymentPaginated PaymentPaginated

	PaymentPaginated.Data = &paymentWithRelation
	PaymentPaginated.Pagination.CurrentPage = int64(params.Page)
	PaymentPaginated.Pagination.Size = int64(params.PageSize)
	PaymentPaginated.Pagination.TotalPages = int64(totalPage)
	PaymentPaginated.Pagination.TotalItem = totalItem

	return nil, &PaymentPaginated, nil

}
