package main

import "discord/internal/controller"

func main() {
	discordBot := controller.NewEndpointHandler("MTIwMjMxOTY5MjcyMzY1ODc3Mw.G2Tz0I.m0t-iwGoWIwu4kfsQEvy3Z4NJV50KNSI_WJubs")
	discordBot.Start()
}
