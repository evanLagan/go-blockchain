# Simple Blockchain Prototype in Go

This is a basic blockchain prototype implemented in Go.  

- Each block contains a timestamp, data, the previous block's hash, and its own hash.
- Blocks are linked together via cryptographic hashes.
- The blockchain is stored in memory and can be printed to the terminal.

## How It Works

- The first block (Genesis Block) is manually created.
- Each new block references the hash of the previous block.
- Hashes are generated using SHA-256 over the block’s data, timestamp, and previous hash.

## Running the Project
go run .  
-> Shows the usage commands
    go run . addblock "Enter some data here"
    go run . printchain
    go run . validate

go run . serve
-> Starts the HTTP server which currently have the follwoing endpoints
    /chain (prints the current blockchain)
    /addblock (adds a block with specified data to the chain)
