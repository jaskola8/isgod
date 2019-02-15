package main

import (
	"fmt"
	"isgod/api"
)

func main() {
	creds := api.Credentials{}
	creds.Username = "michalb2"
	creds.ApiKey = "7ue5BSPJH5uGcdBPiveSHw"
	ann, _ := api.FetchHeaders(creds, 0, 1)
	first := ann[0]
	fmt.Printf("Hash: %s \nTemat: %s \nData: %s, Typ: %d", first.Hash, first.Subject, first.ModifiedDate, first.Type)
}
