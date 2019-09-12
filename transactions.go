package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	//	"math/big"

	"github.com/ethereum/go-ethereum/common"
	//	"github.com/ethereum/go-ethereum/core/types"
	//	"github.com/ethereum/go-ethereum/ethclient"
	//"encoding/json"
)

func monitorTransaction(loctransactionHash string) {
	fmt.Println("loctransactionHash %s:", loctransactionHash)
	txHash := common.HexToHash(loctransactionHash)

	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("@debug tx.Hash().Hex(): %s", tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2

	fmt.Println("@deb isPending %s", isPending) // false

	fmt.Println(prettyPrint(tx))
	//os.Stdout.Write(b)
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
	return "asasd"
}
