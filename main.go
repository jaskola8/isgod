package main

import (
	"github.com/therecipe/qt/widgets"
	"log"
)

var CONFIG Config
var CONFIGFILE string
var APP *widgets.QApplication
var STOP chan struct{}

func main() {
	var err error
	CONFIGFILE = "config.json"
	CONFIG, err = ReadConfig(CONFIGFILE)
	if err != nil {
		log.Fatalf("Couldn't read config file: %v", err)
		return
	}
	STOP = make(chan struct{})
	ExecApp()

}
