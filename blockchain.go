package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Blockchain struct {
	Blocks []Block
}

// The following file and two functions (SaveToDisk() and LoadDisk() ) enable persistence.
const chainFile = "chain.json"

func (bc *Blockchain) SaveToDisk() error {
	data, err := json.MarshalIndent(bc.Blocks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(chainFile, data, 0644)
}

func LoadFromDisk() *Blockchain {
	var blocks []Block

	file, err := os.ReadFile(chainFile)
	if err != nil {
		fmt.Println("No exisiting chain found. Creating a new chain...")
		return NewBlockchain()
	}

	err = json.Unmarshal(file, &blocks)
	if err != nil {
		fmt.Println("Failed to load chain. Starting fresh.")
		return NewBlockchain()
	}

	return &Blockchain{Blocks: blocks}
}

func NewBlockchain() *Blockchain {
	genesisBlock := Block{
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      "Genesis Block",
		PrevHash:  "",
	}
	genesisBlock.MineBlock()

	return &Blockchain{[]Block{genesisBlock}}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := Block{
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.MineBlock()

	if isValidBlock(newBlock, prevBlock) {
		bc.Blocks = append(bc.Blocks, newBlock)
		bc.SaveToDisk()
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
