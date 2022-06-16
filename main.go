package main

// Create block struct
type Block struct {
	Hash         []byte //this means it contains a byte stream / stream of sinlge byte characters
	Content      []byte
	PrevviosHash []byte
}

//Create a chain struct (a collection of blocks)
type BlockChain struct {
	blocks []*Block
}

//function to hash the data in the block
func (b *Block) GetHash() {
info:
}
