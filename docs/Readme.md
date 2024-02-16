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
실행은 정상적으로 이루어졌지만 전송이 되지 않음

## transfer_token/main.go
토큰이 없어서 실행을 못함.

