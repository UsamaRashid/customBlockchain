package main

import (
	blockchain "customblockchain/blockchain"
	"fmt"
)

func main() {

	fmt.Println("Custom Blockchain!")

	var chainHead *blockchain.Block
	genesis := blockchain.BlockData{Transactions: []string{"S2E", "S2Z"}}
	chainHead = blockchain.InsertBlock(genesis, chainHead)
	secondBlock := blockchain.BlockData{Transactions: []string{"E2Alice", "E2Bob", "S2John"}}
	chainHead = blockchain.InsertBlock(secondBlock, chainHead)

	blockchain.ListBlocks(chainHead)
	// blockchain.ChangeBlock("S2E", "S2Trudy", chainHead)
	blockchain.ListBlocks(chainHead)
	blockchain.VerifyChain(chainHead)
}
