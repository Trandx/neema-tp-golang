package service

import (
	. "neema.co.za/rest/utils/models"
	. "neema.co.za/rest/utils/models/_base"
)

func (s *Service) GetPayments(params *Params) (*PaymentWithRelation, *PaymentPaginated, error) {
	// Business logic (if any)
	return s.Repository.GetPayments(params)
}

// func (s *Service) CreatePayment(Payment *PaymentWithRelation) (*PaymentWithRelation, error) {
// 	return s.Repository.CreatePayment(Payment)
// }
