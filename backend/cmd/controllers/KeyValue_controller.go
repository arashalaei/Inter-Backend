package controllers

import (
	"backend-inter/api"
	"backend-inter/cmd/responses"
	"backend-inter/cmd/utils"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
)

func SetKeyValue(client *ethclient.Client, contract *api.Api) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body responses.KeyVlaueBody
		err := c.Bind(&body)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		auth := utils.GetAccountAuth(client, utils.GoDotEnvVariable("PRIVATE_KEY"))
		_, err = contract.Set(auth, big.NewInt(int64(body.Key)), big.NewInt(int64(body.Value)))
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		return c.JSON(
			http.StatusCreated,
			responses.KeyValueResponse{
				Status:  http.StatusCreated,
				Message: "success",
				Data:    &echo.Map{"key": body.Key, "Value": body.Value}})
	}
}

func GetKeyValue(client *ethclient.Client, contract *api.Api) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.QueryParam("key")
		i, err := strconv.Atoi(key)
		if err != nil {
			panic(err)
		}
		value, err := contract.Get(nil, big.NewInt(int64(i)))
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		return c.JSON(
			http.StatusCreated,
			responses.KeyValueResponse{
				Status:  http.StatusOK,
				Message: "success",
				Data:    &echo.Map{"key": key, "value": value}})
	}
}
