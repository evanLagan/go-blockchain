# Simple Blockchain Prototype in Go

This is a basic blockchain prototype implemented in Go.  
It demonstrates the core concepts of a blockchain, including:

- Each block contains a timestamp, data, the previous block's hash, and its own hash.
- Blocks are linked together via cryptographic hashes, ensuring tamper-evidence.
- The blockchain is stored in memory and can be printed to the terminal.

## How It Works

- The first block (Genesis Block) is manually created.
- Each new block references the hash of the previous block.
- Hashes are generated using SHA-256 over the block’s data, timestamp, and previous hash.

## Running the Project
go run .
