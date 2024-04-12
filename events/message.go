package events

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const prefix string = "£"

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == session.State.SessionID {
		return
	}

	// channel, _ := session.Channel(message.ChannelID)
	// guild, _ := session.Guild(message.GuildID)

	args := strings.Split(message.Content, " ")
	if strings.Contains(args[0], prefix) {
		fmt.Println(args)
		command := strings.Trim(args[0], prefix)

		if command == "ping" {
			session.ChannelMessageSend(message.ChannelID, "pong")
		}
	}
}
