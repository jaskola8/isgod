package api

import (
	"log"
	"net/http"
)

const rawApiURL = "https://isod.ee.pw.edu.pl/isod-portal/wapi?q=&username=&apikey="

func CheckConnection() (bool, string) {
	success := false
	client := &http.Client{}
	resp, err := client.Get(rawApiURL)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return false, "HTTP protocol error"
	}
	if resp.StatusCode <= 400 {
		success = true
	}
	return success, resp.Status
}
