package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	. "neema.co.za/rest/utils/logger"
	. "neema.co.za/rest/utils/models"
	. "neema.co.za/rest/utils/models/_base"
)

func (this *Api) DeleteCustomer(c *fiber.Ctx) error {
	var data Customer

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		Logger.Error(err.Error())
		return err
	}

	data.ID = int64(id)

	customer, err := this.Service.DeleteCustomer(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(customer)

}

func (this *Api) CreateCustomer(c *fiber.Ctx) error {
	var data Customer
	if err := c.BodyParser(&data); err != nil {
		Logger.Error(err.Error())
		return err
	}

	customer, err := this.Service.CreateCustomer(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(customer)

}

func (this *Api) GetCustomers(c *fiber.Ctx) error {
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

	customer, customerPaginated, err := this.Service.GetCustomers(&params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if customerPaginated != nil {
		return c.JSON(customerPaginated)
	}

	return c.JSON(customer)

}
