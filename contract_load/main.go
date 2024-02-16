package main

import (
	"fmt"
	"log"
	"os"

	store "ethereum_sdk_demo/contracts" // for demo
	"ethereum_sdk_demo/dot_env"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	dot_env.InitEnv()

	client, err := ethclient.Dial(os.Getenv("ETHEREUM_NETWORK"))
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x7F14DC1E7813d6321AbC2253254c8c1F16146dfa")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}
