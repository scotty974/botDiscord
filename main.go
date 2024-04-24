package main

import (
	"fmt"

	"github.com/SteakBarbare/GoBot-Cours/bot"
	"github.com/SteakBarbare/GoBot-Cours/config"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})

}
