package service

import (
	. "neema.co.za/rest/utils/models"
	. "neema.co.za/rest/utils/models/_base"
)

func (s *Service) GetInvoices(params *Params) (*InvoiceWithRelation, *InvoicePaginated, error) {
	// Business logic (if any)
	return s.Repository.GetInvoices(params)
}

func (s *Service) CreateInvoice(Invoice *InvoiceWithRelation) (*InvoiceWithRelation, error) {
	return s.Repository.CreateInvoice(Invoice)
}

func (s *Service) DeleteInvoice(Invoice *Invoice) (string, error) {
	return s.Repository.DeleteInvoice(Invoice)
}
