package service

import (
	. "neema.co.za/rest/utils/models"
)

func (s *Service) GetCustomerByID(id int) (*Customer, error) {
	// Business logic (if any)
	return s.Repository.GetCustomerByID(id)
}
