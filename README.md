# Blockchain Prototype in Go

This is a basic blockchain prototype implemented in Go.  

- Each block contains a timestamp, data, the previous block's hash, and its own hash.
- Blocks are linked together via cryptographic hashes.
- The blockchain is stored in memory and can be printed to the terminal.
- New blocks can be added over HTTP, this data will then be broadcasted to the other nodes.

## How It Works

- The first block (Genesis Block) is manually created.
- Each new block references the hash of the previous block.
- Hashes are generated using SHA-256 over the block’s data, timestamp, and previous hash.

## Running the Project
```bash
go run .
```
-> Shows the usage commands
    go run . addblock "Enter some data here"
    go run . printchain
    go run . validate
```bash
go run . serve
```
-> Starts the HTTP server which currently have the follwoing endpoints
    /chain (prints the current blockchain)
    /addblock (adds a block with specified data to the chain)

# Demonstration example:

Run the following in three seperate terminals

Node A:
```bash
$env:PORT="8080"
$env:PEERS="http://localhost:8081,http://localhost:8082"
go run . serve
```

Node B (new window):
```bash
$env:PORT="8081"
$env:PEERS="http://localhost:8080,http://localhost:8082"
go run . serve
```

Node C (new window):
```bash
$env:PORT="8082"
$env:PEERS="http://localhost:8080,http://localhost:8081"
go run . serve
```


To view the contents of the chain use:
```bash
curl.exe -s http://localhost:8080/chain
```
Now add a block to any one of the nodes, for example Node B:
```bash
curl.exe -s http://localhost:8081/addblock -X POST -d "Lisa gave Norman 15"
```
Now check the nodes to see that the new block has been successfully broadcasted:
```bash
curl.exe -s http://localhost:8082/chain
curl.exe -s http://localhost:8081/chain
curl.exe -s http://localhost:8080/chain
```
