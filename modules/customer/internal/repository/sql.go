package repository

import (
	. "neema.co.za/rest/utils/logger"
	. "neema.co.za/rest/utils/models"
	. "neema.co.za/rest/utils/pageanable"
)

func (r *Repository) GetCustomerByID(id int) (*Customer, error) {

	var customer []Customer
	err := r.Limit().Find(&customer, &Customer{ID: id})

	if err != nil {
		Logger.Error(err.Error())
		return nil, err
	}

	if len(customer) == 0 {
		return nil, nil
	}
	return &customer[0], nil

}
