package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	. "neema.co.za/rest/utils/logger"
	. "neema.co.za/rest/utils/models"
	. "neema.co.za/rest/utils/models/_base"
)

func (this *Api) DeleteInvoice(c *fiber.Ctx) error {
	var data Invoice

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		Logger.Error(err.Error())
		return err
	}

	data.ID = int64(id)

	customer, err := this.Service.DeleteInvoice(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(customer)

}

func (this *Api) CreateInvoice(c *fiber.Ctx) error {
	var data InvoiceWithRelation
	if err := c.BodyParser(&data); err != nil {
		Logger.Error(err.Error())
		return err
	}

	invoice, err := this.Service.CreateInvoice(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(invoice)

}

func (this *Api) GetInvoices(c *fiber.Ctx) error {
	var params Params

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		id = 0
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		pageSize = 50 // will be put inside the defaulth struct params
	}

	params.PageSize = pageSize
	params.Page = page
	params.Id = int64(id)

	invoice, invoicePaginated, err := this.Service.GetInvoices(&params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if invoicePaginated != nil {
		return c.JSON(invoicePaginated)
	}

	return c.JSON(invoice)

}
