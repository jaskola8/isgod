package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestCreateURL(t *testing.T) {
	outputFull := createURL(Credentials{Username: "u", ApiKey: "a"}, 0, 1, true)
	outputHeaders := createURL(Credentials{Username: "u", ApiKey: "a"}, 0, 1, false)
	expectedFull := "https://isod.ee.pw.edu.pl/isod-portal/wapi?apikey=a&from=0&q=mynewsfull&to=1&username=u"
	expectedHeaders := "https://isod.ee.pw.edu.pl/isod-portal/wapi?apikey=a&from=0&q=mynewsheaders&to=1&username=u"
	if outputFull != expectedFull {
		t.Fatalf("Expected %s, got %s.", expectedFull, outputFull)
	}
	if outputHeaders != expectedHeaders {
		t.Fatalf("Expected %s, got %s.", expectedHeaders, outputHeaders)
	}
}

//Tests fetch function by comparing its output from querying temporary local server with predefined values
func TestFetch(t *testing.T) {
	file, err := os.Open("test_reply.json")
	if err != nil {
		t.Fatalf("Error opening file containing testing data: %v", err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Error reading file containing testing data: %v", err)
	}

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", content)
	})

	go http.ListenAndServe(":3000", nil)

	response, err := fetch("http://localhost:3000/test")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", response.FirstName)
	compareResponse(&response, t)
}

func compareResponse(response *Response, t *testing.T) {
	if response.FirstName != "FirstName" {
		t.Errorf("Expected %s, got %s.", "FirstName", response.FirstName)
	}
	if response.SecondName != "SecondName" {
		t.Errorf("Expected %s, got %s.", "SecondName", response.SecondName)
	}
	if response.Semester != "Semester" {
		t.Errorf("Expected %s, got %s.", "Semester", response.Semester)
	}
	compareAnnoucements(response.Items, t)
}

func compareAnnoucements(annoucements []Annoucement, t *testing.T) {
	if annoucements[0].Content != "Content" {
		t.Errorf("Expected %s, got %s.", "Content", annoucements[0].Content)
	}
	if annoucements[0].Type != 1001 {
		t.Errorf("Expected %d, got %d.", 1001, annoucements[0].Type)
	}
	if annoucements[1].Type != 1002 {
		t.Errorf("Expected %d, got %d.", 1001, annoucements[1].Type)
	}
	if annoucements[1].Content != "" {
		t.Errorf("Expected empty string, got %s", annoucements[1].Content)
	}
}
