package api

import (
	"encoding/json"
	"log"
	"os"
)

type Credentials struct {
	Username string
	APIKey   string
}

func GetCredentials(file *os.File) (Credentials, error) {
	decoder := json.NewDecoder(file)
	credentials := Credentials{}
	err := decoder.Decode(&credentials)
	if err != nil {
		log.Fatal(err)
	}
	return credentials, err
}

func SetCredentials(username, key string, file *os.File) error {
	credentials := map[string]string{"Username": username, "APIKey": key}
	encoder := json.NewEncoder(file)
	err := encoder.Encode(credentials)
	if err != nil {
		log.Fatalf("Error encoding json credentials: %s", err)
	}
	return err
}
