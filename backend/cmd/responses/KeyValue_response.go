package responses

import (
	"github.com/labstack/echo/v4"
)

type KeyValueResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    *echo.Map `json:"data"`
}

type KeyVlaueBody struct {
	Key   int `json:"key"`
	Value int `json:"value"`
}
