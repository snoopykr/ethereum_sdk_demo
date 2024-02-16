package main

import (
	"context"
	"ethereum_sdk_demo/dot_env"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	dot_env.InitEnv()

	client, err := ethclient.Dial(os.Getenv("ETHEREUM_NETWORK"))
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Last Block Number
	fmt.Println(header.Number.String())

	blockNumber := big.NewInt(5297822)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // Block Number
	fmt.Println(block.Time())                // Timestamp
	fmt.Println(block.Difficulty().Uint64()) // Nonce 0
	fmt.Println(block.Hash().Hex())          // Block Hash
	fmt.Println(len(block.Transactions()))   // Transaction Count

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // Transaction Count
}
