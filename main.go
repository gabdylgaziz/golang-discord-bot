package main

import (
	"discord/internal/controller"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	discordBot := controller.NewEndpointHandler(os.Getenv("DiscordToken"))
	discordBot.Start()
}
