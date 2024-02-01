package main

import (
	"discord/internal/controller"
	"os"
)

func main() {
	discordBot := controller.NewEndpointHandler(os.Getenv("DiscordToken"))
	discordBot.Start()
}
