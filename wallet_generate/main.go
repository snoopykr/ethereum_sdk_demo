package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Private Key 생성.
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	// Public Key 생성
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	// Address
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	// Address
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))

	// ffaf0d89b58ddb2955402f819cbe542f6fb0752b468b9ad78f753ddcb5aa1245
	// 08808dde7f1d51bd051f62b3b5e0fb941cae06a762a5f0652f2243e8cca12a9ea074df660becf178530a459a60a18acfc87f3618e1709a731bec787feb60b7d7
	// 0x771Dd54636f0cD9A2642aA128f3020e5cd7F0e30
	// 0x771dd54636f0cd9a2642aa128f3020e5cd7f0e30

}
