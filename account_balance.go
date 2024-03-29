package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func accountBalance(locaddress string) {
	fmt.Println("logging account balance")
	account := common.HexToAddress(locaddress)
	fmt.Println("Account %v", account)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		fmt.Println("error %v", err)
	}
	fmt.Println("my %s account balance : %s", locaddress, balance) // 25893180161173005034

	// blockNumber := big.NewInt(0)
	// balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("balanceAt : %v", balanceAt) // 25729324269165216042

	// fbalance := new(big.Float)
	// fbalance.SetString(balanceAt.String())
	// ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	// fmt.Println("ethValue : %v", ethValue) // 25.729324269165216041

	// pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	// fmt.Println("pendingBalance : %v", pendingBalance) // 25729324269165216042
}
