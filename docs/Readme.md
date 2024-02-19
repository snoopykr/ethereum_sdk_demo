[//]: # (참조 링크)
[//]: # (링크 : [Github][github])

[//]: # ()
[//]: # ([github]: https://github.com/miguelmota/ethereum-development-with-go-book/tree/master "Go Github")

[//]: # (외부 링크)
[//]: # ([Google]&#40;https://google.com, "google link"&#41;)

# Ethereum Development with Go
링크 : <https://github.com/miguelmota/ethereum-development-with-go-book/tree/master>

## transactions/main.go
해결 완료

<strike>transactions는 버전 문제로 인해 실행이 되지 않음</strike>

- (AsIs) github.com/ethereum/go-ethereum v1.13.12
- (ToBe) github.com/ethereum/go-ethereum v1.8.20

v1.8.20 적용후 발생 에러
```
snoopy_kr@MacBookPro ethereum_sdk_demo % go mod tidy
ethereum_sdk_demo/wallet_generate imports
        github.com/ethereum/go-ethereum/crypto imports
        github.com/btcsuite/btcd/btcec: module github.com/btcsuite/btcd@latest found (v0.24.0), but does not contain package github.com/btcsuite/btcd/btcec
```

## transfer_eth/main.go
<strike>실행은 정상적으로 이루어졌지만 전송이 되지 않음</strike>

## transfer_token/main.go
```
0xa9059cbb
0x000000000000000000000000fa9d10ae2052b5e23ae29304f326a625c5c03ee3
0x00000000000000000000000000000000000000000000003635c9adc5dea00000
2024/02/16 15:20:13 execution reverted
```


참조 : <https://medium.com/@blocktorch/shining-light-on-web3-engineering-execution-reverted-errors-cda02f7ae75e>

## transaction_raw_create/main.go
실행은 정상적으로 이루어졌지만 전송이 되지 않음

```
[0xc0000aede0]
```

컴파일 전 에러
```go
	// transfer_eth와 이부분만 틀림...
	// 하지만 전송이 되지 않는다는 거...
	ts := types.Transactions{signedTx}
	fmt.Println(ts)
	
	// 에러 발생 부분.
	// rawTxBytes := ts.GetRlp(0)
	// rawTxHex := hex.EncodeToString(rawTxBytes)
	//
	// fmt.Printf(rawTxHex) // f86...772
```

## transaction_raw_send/main.go
실행은 정상적으로 이루어졌지만 전송이 되지 않음

```
2024/02/16 15:42:41 invalid sender
```

## contract_deploy/main.go
에러 처리 : ^0.4.24 => ^0.8.23
```
snoopy_kr@MacBookPro ethereum_sdk_demo % solc --abi ./contracts/Store.sol
Error: Source file requires different compiler version (current compiler is 0.8.23+commit.f704f362.Darwin.appleclang) - note that nightly builds are considered to be strictly less than the released version
 --> Store.sol:1:1:
  |
1 | pragma solidity ^0.4.24;
  | ^^^^^^^^^^^^^^^^^^^^^^^^
```

에러 처리 : string _version => string memory _version
```
snoopy_kr@MacBookPro ethereum_sdk_demo % solc --abi ./contracts/Store.sol
Warning: SPDX license identifier not provided in source file. Before publishing, consider adding a comment containing "SPDX-License-Identifier: <SPDX-License>" to each source file. Use "SPDX-License-Identifier: UNLICENSED" for non-open-source code. Please see https://spdx.org for more information.
--> Store.sol

Error: Data location must be "storage" or "memory" for constructor parameter, but none was given.
 --> Store.sol:9:15:
  |
9 |   constructor(string _version) public {
  |               ^^^^^^^^^^^^^^^
```

```bash
snoopy_kr@MacBookPro ethereum_sdk_demo % solc --abi ./contracts/Store.sol -o ./contracts
snoopy_kr@MacBookPro ethereum_sdk_demo % solc --bin ./contracts/Store.sol -o ./contracts
snoopy_kr@MacBookPro ethereum_sdk_demo % abigen --bin ./contracts/Store.bin --abi ./contracts/Store.abi --pkg store --out ./contracts/Store.go
```

Etherscan에서 Gas 부족 에러
```
Warning! Error encountered during contract execution [contract creation code storage out of gas]
```

에러 처리 : 30000 => 3000000 (X 100)

Etherscan : <https://sepolia.etherscan.io/tx/0xd0cb57e9e0f8d35300f492034fef3df9d5794731ce489d251ecf1cf863441367>
```
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
```

## contract_read_erc20/main.go

ERC20
```
snoopy_kr@MacBookPro ethereum_sdk_demo % solc --abi ./contracts_erc20/ERC20.sol -o ./contracts_erc20
snoopy_kr@MacBookPro ethereum_sdk_demo % solc --bin ./contracts_erc20/ERC20.sol -o ./contracts_erc20
snoopy_kr@MacBookPro ethereum_sdk_demo % abigen --bin ./contracts_erc20/ERC20.bin --abi ./contracts_erc20/ERC20.abi --pkg token --out ./contracts_erc20/erc20.go
```

## event_subscribe/main.go

```
2024/02/19 14:46:08 notifications not supported
```

## event_read/main.go
version 차이 에러...
```go
        // Assignment count mismatch: 1 = 2
		err := contractAbi.Unpack(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
```

## event_read_erc20/main.go
version 차이 에러...
```go
        // Assignment count mismatch: 1 = 2
		err := contractAbi.Unpack(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
```


## event_read_0xprotocol/main.go
version 차이 에러...
```go
        // Assignment count mismatch: 1 = 2
		err := contractAbi.Unpack(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
```