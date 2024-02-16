package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"os"

	"ethereum_sdk_demo/dot_env"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	dot_env.InitEnv()

	client, err := ethclient.Dial(os.Getenv("ETHEREUM_NETWORK"))
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(5194838)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println("Tx Hash : ", tx.Hash().Hex())             // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println("Tx Value : ", tx.Value().String())        // 10000000000000000
		fmt.Println("Tx Gas : ", tx.Gas())                     // 105000
		fmt.Println("Tx Gas Price : ", tx.GasPrice().Uint64()) // 102000000000
		fmt.Println("Tx Nonce : ", tx.Nonce())                 // 110644
		fmt.Println("Tx Data : ", tx.Data())                   // []
		fmt.Println("Tx To : ", tx.To().Hex())                 // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		// v1.13.12 변경.
		if from, err := types.Sender(types.NewLondonSigner(chainID), tx); err == nil {
			fmt.Println("Tx From : ", from)
		}

		//if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
		//	fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		//}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Tx Status : ", receipt.Status, "\n") // 1
	}

	blockHash := common.HexToHash("0x1d936c2f32914fb7e4ba69000a2d1e4e1b4db458cf9df9f0f1c2b3b30aa8489b")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\n\n\nTx Hash", tx.Hash().Hex()) // Tx Hash
	}

	txHash := common.HexToHash("0x1241ad7d14b6693fb3948a431902f484846794a0a38693147c7a07ac025c6c91")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n\n\nTx Hash", tx.Hash().Hex()) // Tx Hash
	fmt.Println("Tx isPending", isPending)        // Pending 여부
}
