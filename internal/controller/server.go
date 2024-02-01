package controller

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

type EndpointHandler struct {
	Dg     *discordgo.Session
	ApiKey string
}

func NewEndpointHandler(apikey string) *EndpointHandler {
	return &EndpointHandler{
		ApiKey: apikey,
	}
}

func (eh *EndpointHandler) Start() {
	dg, err := discordgo.New("Bot " + eh.ApiKey)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	eh.Dg = dg

	eh.Dg.Identify.Intents = discordgo.IntentsGuildMessages

	eh.AddHandlers()

	err = eh.Dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	defer eh.Dg.Close()

	fmt.Println("Bot is online.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
