package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const token string = "{*some token*}"

var BotID string

var day_people = [5][]string{
	{"Alicia", "Alyssa", "John", "Oli"},
	{"Alicia", "Alyssa", "Arjun", "John", "Oli"},
	{"Alicia", "Alyssa", "Arjun", "John"},
	{"Alicia", "Alyssa", "John", "Oli", "Sanam"},
	{"Alicia", "Alyssa", "Arjun", "Oli", "Sanam"},
}

func main() {

	var err error

	s, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := s.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	s.AddHandler(messageHandler)

	err = s.Open()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Bot is running!")

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Content == "/today" {
		today := time.Now().Weekday()

		var all_people string

		for _, str := range day_people[int(today-1)] {
			all_people += str + ", "
		}

		all_people = strings.TrimSuffix(all_people, ", ")

		_, _ = s.ChannelMessageSend(m.ChannelID, "Who will be in the office today:\n"+today.String()+":\n"+all_people)
	}

	if m.Author.ID == BotID {
		return
	}
}
