package main

import (
	"fmt"
	"os"
	"strings"
)

var chain *Blockchain // global to share between CLI and HTTP

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("	addblock \"some data\"  - Add a new block")
		fmt.Println("	printchain			 	- Show all blocks")
		fmt.Println("	validate				- Check chain validity")
		fmt.Println("	serve					- Run HTTP API (port 8080)")
		return
	}

	command := os.Args[1]

	if command == "serve" {
		chain = LoadFromDisk()
		StartServer("8080")
		return
	}

	switch command {
	case "addblock", "printchain", "validate":
		chain := LoadFromDisk()
		switch command {
		case "addblock":
			if len(os.Args) < 3 {
				fmt.Println("Missing data for new block")
				return
			}
			data := strings.Join(os.Args[2:], " ")
			chain.AddBlock(data)

		case "printchain":
			for i, block := range chain.Blocks {
				fmt.Printf("\n Block #%d\n", i)
				fmt.Printf("Timestamp:    %s\n", block.Timestamp)
				fmt.Printf("Data:         %s\n", block.Data)
				fmt.Printf("PrevHash:     %s\n", block.PrevHash)
				fmt.Printf("Hash:         %s\n", block.Hash)
				fmt.Printf("Nonce:        %d\n", block.Nonce)
			}

		case "validate":
			valid := true
			for i := 1; i < len(chain.Blocks); i++ {
				if !isValidBlock(chain.Blocks[i], chain.Blocks[i-1]) {
					fmt.Printf("Block #%d is invalid!\n", i)
					valid = false
				}
			}
			if valid {
				fmt.Println("Blockchain is valid.")
			}
		}

	default:
		fmt.Println("Unknown command:", command)
	}

}
