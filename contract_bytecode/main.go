package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"os"

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

	contractAddress := common.HexToAddress("0x7F14DC1E7813d6321AbC2253254c8c1F16146dfa")
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode)) // 60806...10029
}
