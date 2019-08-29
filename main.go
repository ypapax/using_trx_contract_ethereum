package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ypapax/jsn"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal("Whoops something went wrong!", err)
	}

	ctx := context.Background()
	tx, pending, err := conn.TransactionByHash(ctx, common.HexToHash("0x30999361906753dbf60f39b32d3c8fadeb07d2c0f1188a32ba1849daac0385a8"))
	if err != nil {
		log.Fatal(err)
	}
	if !pending {
		log.Println(jsn.B(tx))
	} else {
		log.Println("it's pending")
	}
}
