package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	RefreshTimeout    string
	Muted             bool
	FetchSize         int
	RecentFingerprint string
	FontSize          int
	WindowWidth       int
	WindowHeight      int
	WindowPosX        int
	WindowPosY        int
	WindowOnTop       bool
	WindowAutoHide    bool
	DarkTheme         bool
	AesKey            string
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
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(conf)
	return err
}
