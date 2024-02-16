package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0xa26d1cba7dc72aa6a15afb36afed64d64a917a62
}

func importKs() {
	file := "./tmp/UTC--2024-02-15T06-42-05.424569000Z--a26d1cba7dc72aa6a15afb36afed64d64a917a62"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0xa26d1cba7dc72aa6a15afb36afed64d64a917a62

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	createKs()
	//importKs()
}
