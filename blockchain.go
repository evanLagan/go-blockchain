package main

import (
	"time"
)

type Blockchain struct {
	Blocks []Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := Block{
		Timestamp: time.Now(),
		Data:      "Genesis Block",
		PrevHash:  "",
	}

	genesisBlock.GenerateHash()

	return &Blockchain{[]Block{genesisBlock}}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := Block{
		Timestamp: time.Now(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.GenerateHash()
	bc.Blocks = append(bc.Blocks, newBlock)
}
