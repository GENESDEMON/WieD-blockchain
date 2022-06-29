package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// Create block struct
type Block struct {
	Hash         []byte //this means it contains a byte stream / stream of sinlge byte characters
	Content      []byte
	PreviousHash []byte
}

//Create a chain struct (a collection of blocks)
type BlockChain struct {
	blocks []*Block
}

//function to hash the data in the block
func (b *Block) GetHash() {
	info := bytes.Join([][]byte{b.Content, b.PreviousHash}, []byte{})
	// Joins previous block's relevant info with the new blocks
	hash := sha256.Sum256(info)
	//Performs the actual hashing algorithm
	b.Hash = hash[:]
	//This slices

}

//function to create blocks
func CreateBlock(content string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(content), previousHash}
	//this is a pointer
	block.GetHash()
	return block
}

//A chain that ties up our block
func (chain *BlockChain) JoinBlock(content string) {
	previousBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(content, previousBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

//Our first block, since we can't create one from the blockchain, cause it always needs a previous block
func WiedRoot() *Block {
	return CreateBlock("WiedRoot", []byte{})
}

//Initialize a block with the root blocks
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{WiedRoot()}}
}

//Let's print stuff baby
func main() {

	chain := InitBlockChain()

	chain.JoinBlock("Block 2 -")
	chain.JoinBlock("Block 3 -")
	chain.JoinBlock("Block 4 -")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash: %x\n", block.PreviousHash)
		fmt.Printf("data: %s\n", block.Content)
		fmt.Printf("hash: %x\n", block.Hash)
	}

}
