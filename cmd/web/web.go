package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/clebs/garve/block"
	"github.com/clebs/garve/comms"
	"github.com/joho/godotenv"

	"github.com/clebs/garve/chain"
	"github.com/davecgh/go-spew/spew"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	http.HandleFunc("/status", displayBlockchain)
	http.HandleFunc("/add", addBlock)

	p := os.Getenv("Port")

	fmt.Printf("Listening on port %s\n", p)
	if err := http.ListenAndServe(":"+p, nil); err != nil {
		fmt.Printf("error on garve web: %v", err)
	}
}

func displayBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(chain.Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func addBlock(w http.ResponseWriter, r *http.Request) {
	var m comms.Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	newBlock, err := block.New(chain.Blockchain.Top(), m)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}
	if chain.Blockchain.CanAppend(newBlock) {
		chain.Blockchain.Append(newBlock)
		spew.Dump(chain.Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
