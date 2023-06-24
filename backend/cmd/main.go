package main

import (
	"backend-inter/cmd/config"
	"backend-inter/cmd/routes"
	"backend-inter/cmd/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// get contract
	client, contract := config.GetContract()

	// routes
	routes.KeyValueRoute(e, client, contract)

	port := utils.GoDotEnvVariable("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
