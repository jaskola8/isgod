package api

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetCredentials(t *testing.T) {
	file, err := os.Open("test_credentials.json")
	if err != nil {
		t.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()
	credentials, err := ReadCredentials(file)
	if err != nil {
		t.Fatalf("Error reading from file: %s", err)
	}
	if credentials.Username != "username" || credentials.ApiKey != "apikey" {
		t.Fatalf("Credentials do not match predefinied values, got: %s", credentials)
	}
}

func TestWriteCredentials(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "tmp_creds.json")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer tmpfile.Close()
	err = WriteCredentials("username", "apikey", tmpfile)
	if err != nil {
		t.Fatalf("Error writing credentials: %s", err)
	}
	tmpfile.Seek(0, 0)
	credentials, err := ReadCredentials(tmpfile)
	if err != nil {
		t.Fatalf("Error reading from file: %s", err)
	}
	if credentials.Username != "username" || credentials.ApiKey != "apikey" {
		t.Fatalf("Credentials do not match predefinied values, got: %s", credentials)
	}
}

func TestSaveCredentials(t *testing.T) {
	creds := Credentials{"username", "apikey"}
	err := SaveCredentials(creds)
	if err != nil {
		t.Fatalf("Error when setting environment variable: %s", err)
	}
	defer os.Unsetenv("isgod")
	str, ok := os.LookupEnv("isgod")
	if !ok {
		t.Fatalf("Credentials not found after saving")
	}

	if str != "username apikey" {
		t.Fatalf("Credentials do not match predefinied values, got %s", str)
	}
}

func TestReadEnvCredentials(t *testing.T) {
	err := os.Setenv("isgod", "username apikey")
	if err != nil {
		t.Fatalf("Error when setting environment varialbe: %s", err)
	}
	defer os.Unsetenv("isgod")

	creds, err := ReadEnvCredentials()
	if err != nil {
		t.Fatalf("Error during credentials read: %s", err)
	}

	if creds.Username != "username" || creds.ApiKey != "apikey" {
		t.Fatalf("Credentials do not match predefinied values, got: %s", creds)
	}

}
