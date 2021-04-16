package blockchain

import (
	"github.com/JoseRodrigues443/miguel-blockchain-golang/block"
	"github.com/JoseRodrigues443/miguel-blockchain-golang/proof"
)

type BlockChain struct {
	blocks []*block.Block
	proof  proof.Proofer
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := block.Create(data, prevBlock.Hash, chain.proof)
	chain.blocks = append(chain.blocks, new)
}

func (chain *BlockChain) Blocks() []*block.Block {
	return chain.blocks
}

func Genesis(proof proof.Proofer) *block.Block {
	return block.Create("Genesis", []byte{}, proof)
}

func InitBlockChain(proof proof.Proofer) *BlockChain {
	return &BlockChain{[]*block.Block{Genesis(proof)}, proof}
}
