package main

import (
	"os"
	"fmt"
	"encoding/json"
)

type Configuration struct {
	FbAppId       string `json:"fbAppId"`
	FbAppSecret   string `json:"fbAppSecret"`
	FbRedirectUri string `json:"fbRedirectUri"`
	GoogleApiUrl  string `json:"googleApiUrl"`
}

func loadConfiguration() *Configuration {
	configuration := Configuration{}
	if file, err := os.Open("config.json"); err == nil {
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(&configuration); err != nil {
			fmt.Println("Json decode error:", err.Error())
		}
	} else {
		fmt.Println("Error while opening a file: ", err.Error())
	}
	return &configuration
}
