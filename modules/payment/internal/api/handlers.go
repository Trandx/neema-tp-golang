package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	. "neema.co.za/rest/utils/models/_base"
)

// func (this *Api) DeletePayment(c *fiber.Ctx) error {
// 	var data Payment

// 	id, err := strconv.Atoi(c.Params("id"))

// 	if err != nil {
// 		Logger.Error(err.Error())
// 		return err
// 	}

// 	data.ID = int64(id)

// 	customer, err := this.Service.DeletePayment(&data)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.JSON(customer)

// }

// func (this *Api) CreatePayment(c *fiber.Ctx) error {
// 	var data PaymentWithRelation
// 	if err := c.BodyParser(&data); err != nil {
// 		Logger.Error(err.Error())
// 		return err
// 	}

// 	Payment, err := this.Service.CreatePayment(&data)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.JSON(Payment)

// }

func (this *Api) GetPayments(c *fiber.Ctx) error {
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

	Payment, PaymentPaginated, err := this.Service.GetPayments(&params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if PaymentPaginated != nil {
		return c.JSON(PaymentPaginated)
	}

	return c.JSON(Payment)

}
