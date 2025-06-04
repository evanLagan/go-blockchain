package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

const difficulty = 4 // Hash must start with "0000". This value can be changed. Lower values = Less computaion. Higher values = Greater computation

type Block struct {
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

func (b *Block) calculateHash() string {
	record := b.Timestamp + b.Data + b.PrevHash + strconv.Itoa(b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// Proof of work
func (b *Block) MineBlock() {
	for {
		hash := b.calculateHash()
		if strings.HasPrefix(hash, strings.Repeat("0", difficulty)) {
			b.Hash = hash
			fmt.Printf("Mined block with nonce %d: %s\n", b.Nonce, hash)
			break
		}
		b.Nonce++
	}
}
