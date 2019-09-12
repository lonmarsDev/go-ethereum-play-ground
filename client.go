package main

import (
	"fmt"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func main() {

	var err error
	client, err = ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	generateWallet()
	value := big.NewInt(2000000000000000000) // in wei (1 eth)
	transferEth("5CD956063F3EB0F959DD961B04142314D78D3E400A504B644A29A96CC38852E6", value)
	accountBalance(address)
	monitorTransaction(transactionHash)

	generateWallet()
	value = big.NewInt(0000000000000100000) // in wei (1 eth)
	transferEth(privateKey2, value)
	accountBalance(address)
	monitorTransaction(transactionHash)
}
