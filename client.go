package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client
func main() {

	var err error
	client, err = ethclient.Dial("https://mainnet.infura.io/v3/7d3b16cd643e480ba8011af9e15bf8ed")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	generateWallet()
	accountBalance()
}
