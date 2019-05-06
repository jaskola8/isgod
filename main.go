package main

import (
	"fmt"
	"isgod/api"
	"os"
)

func main() {
	file, err := os.Open("creds.json")
	if err != nil {
		os.Exit(1)
	}
	creds, err := api.ReadCredentials(file)
	if err != nil {
		os.Exit(2)
	}
	resp, err := api.FetchHeaders(creds, 0, 1)
	if err != nil {
		os.Exit(3)
	}
	ann := resp.Items
	first := ann[0]
	fmt.Printf("Hash: %s \nTemat: %s \nData: %s, Typ: %d", first.Hash, first.Subject, first.ModifiedDate, first.Type)
}
