package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	prefix = "!"
)

func main() {
	dg, err := discordgo.New("Bot " + "MTIwMjMxOTY5MjcyMzY1ODc3Mw.G2Tz0I.m0t-iwGoWIwu4kfsQEvy3Z4NJV50KNSI_WJubs")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	defer dg.Close()

	fmt.Println("Bot is online.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Для того чтобы бот не отвечал сам себе
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, prefix) {
		command, args := parseCommand(m.Content)

		fmt.Println(args)

		if command == "hello" {
			s.ChannelMessageSend(m.ChannelID, "Hello, "+m.Author.Username+"! (Asynchronous)")
		}
	}

	//fmt.Println(m.Content)
}

func parseCommand(message string) (command string, args []string) {
	parts := strings.Fields(message[len(prefix):])
	if len(parts) > 0 {
		command = strings.ToLower(parts[0])
		args = parts[1:]
	}

	return
}
