package main

import (
	"fmt"
)

func main() {
	chain := NewBlockchain()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")

	for i, block := range chain.Blocks {
		fmt.Printf("Block #%d\n", i)
		fmt.Printf("Timestamp: %s\n", block.Timestamp.String())
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Previous Hash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}
}
