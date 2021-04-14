package main

import (
	"fmt"

	"github.com/JoseRodrigues443/miguel-blockchain-golang/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	fmt.Printf("Genesis is done\n")

	chain.AddBlock("Block One")
	chain.AddBlock("Block Two")
	chain.AddBlock("Block Three")

	for index, block := range chain.Blocks() {
		fmt.Printf("______________\nBlock number ==>  %x\n", index)
		fmt.Printf("Previous ==>  %x\n", block.PrevHash)
		fmt.Printf("Data ==>  %s\n", block.Data)
		fmt.Printf("Hash ==> %x\n", block.Hash)
	}
}
