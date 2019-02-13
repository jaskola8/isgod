//Provides simple api connection check, quering and credentials management
package api

import (
	"log"
	"net/http"
)

const rawApiURL = "https://isod.ee.pw.edu.pl/isod-portal/wapi?q=?&username=?&apikey=?"

//Checks if API is accessible, accepts up to 10 redirects, returns boolean and HTTP response code.
//In case of too many redirects or HTTP protocol error logs returns false and "HTTP protocol error"
//as response code.
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
