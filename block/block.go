package block

import "github.com/JoseRodrigues443/miguel-blockchain-golang/proof"

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func Create(data string, prevHash []byte, proof proof.Proofer) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	nonce, hash := proof.Run(block.Data, block.Hash, block.PrevHash, block.Nonce)

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func Copy(Hash []byte, Data []byte, PrevHash []byte, Nonce int) *Block {
	block := &Block{Hash: Hash, Data: Data, PrevHash: PrevHash, Nonce: Nonce}
	return block
}
