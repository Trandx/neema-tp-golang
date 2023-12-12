package repository

import (
	"math"
	"math/rand"
	"strconv"
	"time"

	. "neema.co.za/rest/utils/logger"
	. "neema.co.za/rest/utils/models"
	. "neema.co.za/rest/utils/models/_base"
)

func randomize() *rand.Rand {
	return rand.New(
		rand.NewSource(time.Now().UnixNano()))
}

func (r *Repository) DeleteCustomer(customer *Customer) (string, error) {

	_, err := r.Delete(customer)

	if err != nil {
		Logger.Error(err.Error())
		return "", err
	}

	Logger.Info("Customer deleted")

	return "Customer deleted successfully", nil
}

func (r *Repository) CreateCustomer(customer *Customer) (*Customer, error) {
	var seededRand = randomize()

	customer.State = "Cameroon"
	customer.Id_country = 237
	customer.Id_currency = 550
	customer.Alias = strconv.FormatInt(randomize().Int63(), 16)
	customer.Tmc_client_number = strconv.FormatInt(randomize().Int63(), 16)
	customer.Account_number = strconv.FormatInt(randomize().Int63(), 16)
	customer.Slug = seededRand.Int63()
	customer.Ab_key = strconv.FormatInt(randomize().Int63(), 16)

	_, err := r.Insert(customer)

	if err != nil {
		Logger.Error(err.Error())
		return nil, err
	}

	Logger.Info("Customer saved")

	return customer, nil
}

func (r *Repository) GetCustomers(params *Params) (*Customer, *CustomerPaginated, error) {
	var customers []Customer

	if params.Id > 0 {

		err := r.Find(&customers, &Customer{ID: params.Id})

		if err != nil {
			Logger.Error(err.Error())
			return nil, nil, err
		}

		if len(customers) == 0 {
			return nil, nil, nil
		}

		return &customers[0], nil, nil
	}

	offset := (params.Page - 1) * params.PageSize
	totalItem, err := r.Count(&Customer{})
	totalPage := math.Ceil(float64(totalItem) / float64(params.PageSize))

	if err != nil {
		return nil, nil, err
	}

	err = r.Limit(params.PageSize, offset).Find(&customers)
	if err != nil {
		return nil, nil, err
	}
	var customerPaginated CustomerPaginated

	customerPaginated.Data = &customers
	customerPaginated.Pagination.CurrentPage = int64(params.Page)
	customerPaginated.Pagination.Size = int64(params.PageSize)
	customerPaginated.Pagination.TotalPages = int64(totalPage)
	customerPaginated.Pagination.TotalItem = totalItem

	return nil, &customerPaginated, nil

}
