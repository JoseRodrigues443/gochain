package main

import (
	"fmt"
	"strconv"

	"github.com/JoseRodrigues443/miguel-blockchain-golang/blockchain"
	"github.com/JoseRodrigues443/miguel-blockchain-golang/proof/ofWork"
)

func main() {
	proofOfWork := ofWork.Init()
	chain := blockchain.InitBlockChain(proofOfWork)
	fmt.Printf("Genesis is done\n")

	chain.AddBlock("Block One")
	chain.AddBlock("Block Two")
	chain.AddBlock("Block Three")

	for index, block := range chain.Blocks() {
		fmt.Printf("______________\nBlock number ==>  %x\n", index)
		fmt.Printf("Previous ==>  %x\n", block.PrevHash)
		fmt.Printf("Data ==>  %s\n", block.Data)
		fmt.Printf("Hash ==> %x\n", block.Hash)

		pow := ofWork.Init()
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate(block.Data, block.Hash, block.PrevHash, block.Nonce)))
		fmt.Println()

	}
}
