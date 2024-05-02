package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MemberJoin(session *discordgo.Session, member *discordgo.GuildMemberAdd) {
	fmt.Printf("Member joined the server: %s\n", member.User.Username)

	channel, err := session.UserChannelCreate(member.User.ID)
	if err != nil {
		fmt.Println("Error creating direct message channel: ", err)
		return
	}

	guild, err := session.Guild(member.GuildID)
	if err != nil {
		fmt.Println("Error getting guild: ", err)
		return
	}

	session.ChannelMessageSend(channel.ID, fmt.Sprintf("Welcome to the %s server!\n"+
		"Please register with the bot using the `/register` command in the server.\n"+
		"To register with the clan please make a ticket in <#1225156608145887243>.", guild.Name))
}
