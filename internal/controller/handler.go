package controller

import (
	"discord/internal/api"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (eh *EndpointHandler) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Для того чтобы бот не отвечал сам себе
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, "!") {
		command, args := parseCommand(m.Content)

		//	сюда можно добавить любую команду, которая будет нужно, args это все что будет написано после команды
		//	например
		// 	!weather Almaty - weather это команда, Almaty это args
		//
		switch command {
		case "help":
			go helpHandler(s, m)
		case "weather":
			go weatherHandler(s, m, args)
		case "translate":
			go translateHandler(s, m, args)
		}
	}

}

func parseCommand(message string) (command string, args []string) {
	parts := strings.Fields(message[len("!"):])
	if len(parts) > 0 {
		command = strings.ToLower(parts[0])
		args = parts[1:]
	}

	return
}

func helpHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Для того чтобы бот не отвечал сам себе
	if m.Author.ID == s.State.User.ID {
		return
	}

	help := fmt.Sprintf("Привет, %s. Вот список доступных команд: \n !weather название города - Посмотреть погоду определенного города \n !translate <язык> <Текст> - перевод текста на определенный язык ", m.Author.Username)

	s.ChannelMessageSend(m.ChannelID, help)
}

func weatherHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	weather := api.GetWeather(args[0])

	answer := fmt.Sprintf(
		"Прогноз погода в городе %s:\n Погода: %s\n °C: %f\n °F: %f\n",
		weather.Name,
		weather.Text,
		weather.TempC,
		weather.TempF,
	)

	s.ChannelMessageSend(m.ChannelID, answer)
}

func translateHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	resultString := strings.Join(args[1:], " ")
	sourceLang := api.Detect(resultString)
	targetLang := args[0]

	translatedText := api.Translate(resultString, targetLang, sourceLang)

	s.ChannelMessageSend(m.ChannelID, translatedText)
}
