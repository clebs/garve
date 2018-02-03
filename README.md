# Garve
Garve is a blockchain implementation in Go. For now it is a very simple web server that displays the chain as JSON and a CLI tool that allows forges requests to add new blocks to the chain.
Each block contains a string as content, any type implementing the `Stringer` interface can be a block.

## Usage
First start the web server (default on port `8080`) and open `http://localhost:8080/status` on your browser to see the chain.
To start the server run the following command from the root of the project:
```bash
go run cmd/web/web.go
```

To add a new block to the chain run the garve CLI tool and give a message as flag:
```bash
go run cmd/garve/garve.go -m "Hello Blockchain"
```

Now refresh your browser to see your message added to the chain! 😁