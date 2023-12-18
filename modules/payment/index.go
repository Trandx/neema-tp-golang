package payment

import (
	. "neema.co.za/rest/modules/payment/internal/api"
)

func GetModule() *Module {
	api := BuildApi()
	handleRoutes(api)
	module := Module(*api) //Module is an alias of Api
	return &module
}

func handleRoutes(api *Api) {
	api.Get("/:id", api.GetPayments)
	api.Get("/", api.GetPayments)
	//api.Post("", api.CreatePayment)
	//api.Delete("/:id", api.DeletePayment)
}
