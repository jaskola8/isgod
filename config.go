package main

import (
	"encoding/json"
	"isgod/api"
	"os"
)

type Config struct {
	Credentials       api.Credentials
	RefreshTimeout    int
	FetchSize         int
	RecentFingerprint string
}

//Parses specified json file into Credentials type
func ReadConfig(filename string) (Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	decoder := json.NewDecoder(file)
	conf := Config{}
	err = decoder.Decode(&conf)
	return conf, err
}

//Writes username and apikey to file in json format
func (conf *Config) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(conf)
	return err
}
