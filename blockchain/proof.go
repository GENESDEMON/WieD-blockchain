package blockchain

import (
	"math/big"
)

const Difficulty = 13

//declaring the struct
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

//initializing the struct
func CreateProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{b, target}
	return pow
}

//
