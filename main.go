// main.go
package main

import (
	"github.com/joho/godotenv"

	customerModule "neema.co.za/rest/modules/customer"
	userModule "neema.co.za/rest/modules/user"
	App "neema.co.za/rest/utils/app"
	. "neema.co.za/rest/utils/logger"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		Logger.Error("Error loading .env file")
	}
}

const API_V1_BASE_PATH = "/api/v1"

func main() {

	app := App.Initialise()
	defer app.Listen(":8080")

	routerV1 := app.Group(API_V1_BASE_PATH)

	routerV1.Mount("/users", userModule.GetModule().App)
	routerV1.Mount("/customers", customerModule.GetModule().App)
}
