package config

import (
	"fmt"
	"os"
)

var (
	TOKEN      string
	BOT_PREFIX string
)

func LoadConfig() error {
	fmt.Println("Loading env values")

	TOKEN = os.Getenv("TOKEN")
	BOT_PREFIX = os.Getenv("BOT_PREFIX")
	fmt.Println("Successfully loaded env variable")
	return nil
}
