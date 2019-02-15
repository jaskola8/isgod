package api

import (
	"encoding/json"
	"log"
	"os"
)

//Type storing username and apikey, used in API authentication
type Credentials struct {
	Username string
	ApiKey   string
}

//Parses specified json file into Credentials type
func GetCredentials(file *os.File) (Credentials, error) {
	decoder := json.NewDecoder(file)
	credentials := Credentials{}
	err := decoder.Decode(&credentials)
	if err != nil {
		log.Fatal(err)
	}
	return credentials, err
}

//Writes username and apikey to file in json format
func WriteCredentials(username, key string, file *os.File) error {
	credentials := map[string]string{"Username": username, "APIKey": key}
	encoder := json.NewEncoder(file)
	err := encoder.Encode(credentials)
	if err != nil {
		log.Fatalf("Error encoding json credentials: %s", err)
	}
	return err
}
