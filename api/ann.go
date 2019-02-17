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
	ApiUrl = "https://isod.ee.pw.edu.pl/isod-portal/wapi"
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

type Result struct {
	Items      []Annoucement
	FirstName  string
	SecondName string
	Semester   string
}

func createURL(creds Credentials, from, to int, full bool) string {
	url, _ := url.Parse(ApiUrl)
	query := url.Query()
	if full {
		query.Set("q", "mynewsfull")
	} else {
		query.Set("q", "mynewsheaders")
	}
	query.Set("username", creds.Username)
	query.Set("apikey", creds.APIKey)
	query.Set("from", strconv.Itoa(from))
	query.Set("to", strconv.Itoa(to))
	url.RawQuery = query.Encode()
	return url.String()
}

func FetchFull(creds Credentials, from, to int) ([]Annoucement, error) {
	data, err := fetch(createURL(creds, from, to, true))
	return data.Items, err
}

func FetchHeaders(creds Credentials, from, to int) ([]Annoucement, error) {
	data, err := fetch(createURL(creds, from, to, false))
	return data.Items, err
}

func fetch(url string) (result Result, err error) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return result, fmt.Errorf("GET method on %s: %v", url, err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("reading query body from %s: %v", url, err)
	}
	json.Unmarshal(data, &result)
	return result, nil
}
