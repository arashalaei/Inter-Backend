package routes

import (
	"backend-inter/api"
	"backend-inter/cmd/controllers"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
)

func KeyValueRoute(e *echo.Echo, client *ethclient.Client, contract *api.Api) {
	e.POST("/api/v1/set", controllers.SetKeyValue(client, contract))
	e.GET("/api/v1/retrieve", controllers.GetKeyValue(client, contract))
}
