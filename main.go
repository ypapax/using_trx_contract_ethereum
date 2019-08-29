package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}

	contract, err := NewTronToken(common.HexToAddress("0xf230b790E05390FC8295F4d3F60332c93BEd42e2"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	amt, err := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x387fc6939b5e54b2f11793df05388f9d11942948"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("balance amount: ", amt)
}
