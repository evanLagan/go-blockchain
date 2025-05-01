package main

import (
	"fmt"
	"strings"
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
	genesisBlock.MineBlock()

	return &Blockchain{[]Block{genesisBlock}}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := Block{
		Timestamp: time.Now(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.MineBlock()

	if isValidBlock(newBlock, prevBlock) {
		bc.Blocks = append(bc.Blocks, newBlock)
		fmt.Println("Block added to the chain")
	} else {
		fmt.Println("Invalid block. Not added")
	}
}

func isValidBlock(newBlock, prevBlock Block) bool {
	if newBlock.PrevHash != prevBlock.Hash {
		return false
	}

	expectedHash := newBlock.calculateHash()
	if newBlock.Hash != expectedHash {
		return false
	}

	// Check proof of work
	if !strings.HasPrefix(newBlock.Hash, strings.Repeat("0", difficulty)) {
		return false
	}

	return true
}
