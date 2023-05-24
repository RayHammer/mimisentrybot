package main

import (
	"log"
	"mimisentry/bot"
	"mimisentry/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	bot.Run()
	return
}
