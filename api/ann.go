package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	ApiURL = "https://isod.ee.pw.edu.pl/isod-portal/wapi"
)

type Annoucement struct {
	Content       string
	Hash          string
	Subject       string
	ModifiedDate  string
	ModifiedBy    string
	Type          int
	NoAttachments int
}

type Response struct {
	Items      []Annoucement
	FirstName  string
	SecondName string
	Semester   string
}

//Creates URL string from given arguments
func createURL(creds Credentials, from, to int, full bool) string {
	url, _ := url.Parse(ApiURL)
	query := url.Query()
	if full {
		query.Set("q", "mynewsfull")
	} else {
		query.Set("q", "mynewsheaders")
	}
	query.Set("username", creds.Username)
	query.Set("apikey", creds.ApiKey)
	query.Set("from", strconv.Itoa(from))
	query.Set("to", strconv.Itoa(to))
	url.RawQuery = query.Encode()
	return url.String()
}

//Query database for annoucements, including their content
func FetchFull(creds Credentials, from, to int) (Response, error) {
	data, err := fetch(createURL(creds, from, to, true))
	return data, err
}

//Query database for annoucements, fetching everything except content
func FetchHeaders(creds Credentials, from, to int) (Response, error) {
	data, err := fetch(createURL(creds, from, to, false))
	return data, err
}

//Returns data returned from GET method on given URL string
func fetch(url string) (result Response, err error) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return result, fmt.Errorf("GET method on %s: %v", url, err)
	}
	if resp.StatusCode != 200 {
		return result, fmt.Errorf("Incorrect response status code: %s", resp.Status)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("reading query body from %s: %v", url, err)
	}
	err = json.Unmarshal(data, &result)
	return result, err
}
