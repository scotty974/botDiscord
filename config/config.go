package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	TOKEN      string
	BOT_PREFIX string
	config     *Config
)

type Config struct {
	TOKEN      string `json:"TOKEN"`
	BOT_PREFIX string `json:"BOT_PREFIX"`
}

func LoadConfig() error {
	fmt.Println("Loading config file")
	file, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Error loading config file")
		return err
	}

	fmt.Println("Converting config file")
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("Error converting config file")
		return err
	}

	TOKEN = config.TOKEN
	BOT_PREFIX = config.BOT_PREFIX
	return nil
}
