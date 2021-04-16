package proof

type Proofer interface {
	Validate(Hash []byte, Data []byte, PrevHash []byte, Nonce int) bool
	Run(Hash []byte, Data []byte, PrevHash []byte, Nonce int) (int, []byte)
}
