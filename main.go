package main

import (
	"discord/internal/controller"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Загружает переменные с .env файла
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// Инициализация бота
func main() {
	discordBot := controller.NewEndpointHandler(os.Getenv("DiscordToken"))
	discordBot.Start()
}
