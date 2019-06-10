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

type Fingerprint struct {
	Generated   string
	Fingerprint string
}

//Creates URL string from given arguments
func createURL(creds Credentials, from, to int, full bool) string {
	url2, _ := url.Parse(ApiURL)
	query := url2.Query()
	if full {
		query.Set("q", "mynewsfull")
	} else {
		query.Set("q", "mynewsheaders")
	}
	query.Set("username", creds.Username)
	query.Set("apikey", creds.ApiKey)
	query.Set("from", strconv.Itoa(from))
	query.Set("to", strconv.Itoa(to))
	url2.RawQuery = query.Encode()
	return url2.String()
}

//Query database for annoucements, fetching everything except content
func FetchAnnoucements(creds Credentials, from, to int, full bool) (Response, error) {
	data, err := fetch(createURL(creds, from, to, full))
	if err != nil {
		return Response{}, err
	}
	var resp Response
	err = json.Unmarshal(data, &resp)
	return resp, err
}

func FetchFingerprint(creds Credentials) (Fingerprint, error) {
	url2, _ := url.Parse(ApiURL)
	query := url2.Query()
	query.Set("q", "mynewsfingerprint")
	query.Set("username", creds.Username)
	query.Set("apikey", creds.ApiKey)
	url2.RawQuery = query.Encode()
	data, err := fetch(url2.String())
	if err != nil {
		return Fingerprint{}, err
	}
	var finger Fingerprint
	err = json.Unmarshal(data, &finger)
	return finger, err
}

//Returns data returned from GET method on given URL string
func fetch(url string) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET method on %s: %v", url, err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Incorrect response status code: %s", resp.Status)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, fmt.Errorf("reading query body from %s: %v", url, err)
	}
	return data, nil
}
