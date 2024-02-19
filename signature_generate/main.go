package main

import (
	"ethereum_sdk_demo/dot_env"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	dot_env.InitEnv()

	privateKey, err := crypto.HexToECDSA(os.Getenv("ETHEREUM_PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0xcd4aeea262135873b49b519a7c7cb7b4b2c610d932a9900f8d3bee88108810a368ebfafa5a2bf634e71d154eaadf3a2fa9e62012a2086ac62d29caf152296fd201
}
