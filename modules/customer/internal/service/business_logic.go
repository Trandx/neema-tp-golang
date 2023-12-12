package service

import (
	. "neema.co.za/rest/utils/models"
	. "neema.co.za/rest/utils/models/_base"
)

func (s *Service) GetCustomers(params *Params) (*Customer, *CustomerPaginated, error) {
	// Business logic (if any)
	return s.Repository.GetCustomers(params)
}

func (s *Service) CreateCustomer(customer *Customer) (*Customer, error) {
	return s.Repository.CreateCustomer(customer)
}

func (s *Service) DeleteCustomer(customer *Customer) (string, error) {
	return s.Repository.DeleteCustomer(customer)
}
