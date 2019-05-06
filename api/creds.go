package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

//Type storing username and apikey, used in API authentication
type Credentials struct {
	Username string
	ApiKey   string
}

//Parses specified json file into Credentials type
func ReadCredentials(file *os.File) (Credentials, error) {
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

//Saves credentials to environmental variable
func SaveCredentials(creds Credentials) error {

	err := os.Setenv("isgod", fmt.Sprintf("%s:%s", creds.Username, creds.ApiKey))
	if err != nil {
		return err
	}
	return nil
}

//Reads credentials to environmental variable
func ReadEnvCredentials() (Credentials, error) {
	str, ok := os.LookupEnv("isgod")
	if !ok {
		return Credentials{}, errors.New("Credentials not found")
	}
	var username, key string
	_, err := fmt.Scanf(str, "%s:%s", username, key)
	if err != nil {
		return Credentials{}, fmt.Errorf("Couldn't parse credentials: %s", err)
	}
	return Credentials{username, key}, nil

}
