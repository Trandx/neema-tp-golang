package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (this *Api) GetCustomerByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	customer, err := this.Service.GetCustomerByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(customer)
}
