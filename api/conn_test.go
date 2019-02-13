package api

import (
	"testing"
)

func TestCheckConnection(t *testing.T) {
	success, response := CheckConnection()
	if !success {
		t.Fatal(response)
	}
}
