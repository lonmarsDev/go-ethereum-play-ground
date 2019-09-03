package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client
func main() {

	var err error
	client, err = ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	generateWallet()
}
