package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Route struct {
	Path    string   `json:"path"`
	Targets []string `json:"targets"`
}

type Config struct {
	Port      int     `json:"port"`
	StaticDir string  `json:"staticDir"`
	Routes    []Route `json:"routes"`
}

var Vars Config

var Retrials int

func GetConfig(path string) {
	configFile, openErr := os.Open(path)
	if openErr != nil {
		fmt.Println("Error opening config file:", openErr)
		panic(openErr)
	}

	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			fmt.Println("Error closing config file:", err)
		}
	}(configFile)

	configBytes, bytesErr := io.ReadAll(configFile)
	if bytesErr != nil {
		fmt.Println("Error opening config file:", openErr)
		panic(openErr)
	}

	jsonErr := json.Unmarshal(configBytes, &Vars)
	if jsonErr != nil {
		fmt.Println("Error parsing config file:", jsonErr)
		panic(jsonErr)
	}

	Retrials = len(Vars.Routes)

	// Cleanups
	cleanStaticDir()
}
