package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	bot "github.com/njh18/tcg-tracker-discord-bot/bot"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	bot.BotToken = os.Getenv("BOT_TOKEN")
	bot.Run() // call the run function of bot/bot.go
}
