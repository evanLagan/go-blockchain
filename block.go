package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Timestamp time.Time
	Data      string
	PrevHash  string
	Hash      string
}

func (b *Block) GenerateHash() {
	record := b.Timestamp.String() + b.Data + b.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	b.Hash = hex.EncodeToString(hashed)
}
