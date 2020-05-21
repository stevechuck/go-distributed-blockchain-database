package main

import (
	"fmt"
	"schuck017/go-distributed-blockchain-database/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First block")
	chain.AddBlock("Second block")
	chain.AddBlock("Third block")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
