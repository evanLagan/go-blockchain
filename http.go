package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func StartServer(port string) {
	http.HandleFunc("/chain", handleChain)
	http.HandleFunc("/addblock", handleAddBlock)
	http.HandleFunc("/receive", handleReceiveBlock)

	fmt.Printf("HTTP server running on http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

func handleChain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chain.Blocks)
}

func handleAddBlock(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	data := string(body)

	chain.AddBlock(data)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Block added\n"))
}

func handleReceiveBlock(w http.ResponseWriter, r *http.Request) {
	var newBlock Block
	err := json.NewDecoder(r.Body).Decode(&newBlock)
	if err != nil {
		http.Error(w, "Invalid block data", http.StatusBadRequest)
		return
	}

	lastBlock := chain.Blocks[len(chain.Blocks)-1]
	if isValidBlock(newBlock, lastBlock) {
		chain.Blocks = append(chain.Blocks, newBlock)
		chain.SaveToDisk()
		w.Write([]byte("Block accepted and added.\n"))
	} else {
		http.Error(w, "Invaid Block", http.StatusBadRequest)
	}

}
