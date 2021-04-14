package blockchain
import (
	"github.com/JoseRodrigues443/miguel-blockchain-golang/block"
)

type BlockChain struct {
	blocks []*block.Block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := block.CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func (chain *BlockChain) Blocks() []*block.Block {
	return chain.blocks
}

func Genesis() *block.Block {
	return block.CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*block.Block{Genesis()}}
}
