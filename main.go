package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
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

	account := common.HexToAddress(os.Getenv("ETHEREUM_FROM_ADDRESS"))

	// 현재 발란스
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	// 지정한 블록넘버 당시 발란스
	blockNumber := big.NewInt(4285686)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt)

	// Eth 단위 변환.
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	// 팬딩 발란스
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance)
}
