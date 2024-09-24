package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	bot "github.com/njh18/tcg-tracker-discord-bot/bot"
	webscraper "github.com/njh18/tcg-tracker-discord-bot/web-scraper"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	bot.BotToken = os.Getenv("BOT_TOKEN")
	// bot.Run()

	webscraper.Main()
}
