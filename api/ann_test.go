package api

import (
	"testing"
)

func TestCreateURL(t *testing.T) {
	outputFull := createURL(Credentials{Username: "u", APIKey: "a"}, 0, 1, true)
	outputHeaders := createURL(Credentials{Username: "u", APIKey: "a"}, 0, 1, false)
	expectedFull := "https://isod.ee.pw.edu.pl/isod-portal/wapi?apikey=a&from=0&q=mynewsfull&to=1&username=u"
	expectedHeaders := "https://isod.ee.pw.edu.pl/isod-portal/wapi?apikey=a&from=0&q=mynewsheaders&to=1&username=u"
	if outputFull != expectedFull {
		t.Fatalf("Expected %s, got %s.", expectedFull, outputFull)
	}
	if outputHeaders != expectedHeaders {
		t.Fatalf("Expected %s, got %s.", expectedHeaders, outputHeaders)
	}
}
