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

func fetch(creds Credentials, from, to int, full bool) (result Result, err error) {
	url, err := url.Parse(ApiUrl)
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

	response, err := Query(url.String())
	if err != nil {
		return result, fmt.Errorf("error fetching data from %s: %v", url.String(), err)
	}
	json.Unmarshal(response, &result)
	return result, nil
}

func FetchFull(creds Credentials, from, to int) ([]Annoucement, error) {
	data, err := fetch(creds, from, to, true)
	return data.Items, err
}

func FetchHeaders(creds Credentials, from, to int) ([]Annoucement, error) {
	data, err := fetch(creds, from, to, false)
	return data.Items, err
}

func Query(url string) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET method on %s: %v", url, err)
	}
	defer resp.Body.Close()
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading query body: %v", url, err)
	}
	return jsonData, nil
}
