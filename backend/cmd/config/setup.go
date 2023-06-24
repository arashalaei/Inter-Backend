package config

import (
	"backend-inter/api"
	"backend-inter/cmd/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetContract() (*ethclient.Client, *api.Api) {

	// address of etherum env
	ANVIL_URL_RPC := utils.GoDotEnvVariable("ANVIL_URL_RPS")
	CONTRACT_ADDRESS := utils.GoDotEnvVariable("CONTRACT_ADDRESS")
	// PRIVATE_KEY := utils.GoDotEnvVariable("PRIVATE_KEY")

	client, err := ethclient.Dial(ANVIL_URL_RPC)
	utils.ErrorHandler(err)

	address := common.HexToAddress(CONTRACT_ADDRESS)
	contract, err := api.NewApi(address, client)
	utils.ErrorHandler(err)

	return client, contract
	// auth := utils.GetAccountAuth(client, PRIVATE_KEY)
	// _, err := contract.Set(auth, big.NewInt(0), big.NewInt(10))
	// g, err := contract.Get(nil, big.NewInt(0))
	// utils.ErrorHandler(err)
}
