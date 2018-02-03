package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/clebs/garve/comms"
	"github.com/joho/godotenv"
)

var message string

// init parses all cli flags and readies the tool for use.
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&message, "message", "", "Message to be added to the blockchain")
	flag.StringVar(&message, "m", "", "(short) Message to be added to the blockchain")

	flag.Parse()
}

// main is the entry point of the garve cmd.
func main() {
	url := fmt.Sprintf("http://localhost:%s/add", os.Getenv("Port"))

	m := comms.Message(message)
	json, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Print("Message sent...")
	fmt.Println(resp.Status)
}
