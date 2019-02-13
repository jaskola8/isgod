package api

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetCredentials(t *testing.T) {
	file, err := os.Open("test_credentials.json")
	defer file.Close()
	if err != nil {
		t.Fatalf("Error opening file: %s", err)
	}
	credentials, err := GetCredentials(file)
	if err != nil {
		t.Fatalf("Error reading from file: %s", err)
	}
	if credentials.Username != "username" || credentials.APIKey != "apikey" {
		t.Fatalf("Wrong credentials read: %s", credentials)
	}
}

func TestWriteCredentials(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "tmp_creds.json")
	defer tmpfile.Close()
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	err = WriteCredentials("username", "apikey", tmpfile)
	if err != nil {
		t.Fatalf("Error writing credentials: %s", err)
	}
	tmpfile.Seek(0, 0)
	credentials, err := GetCredentials(tmpfile)
	if err != nil {
		t.Fatalf("Error reading from file: %s", err)
	}
	if credentials.Username != "username" || credentials.APIKey != "apikey" {
		t.Fatalf("Wrong credentials read: %s", credentials)
	}
}
