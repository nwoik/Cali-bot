package events

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	m "github.com/nwoik/calibotapi/member"
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
		} else if command == "members" {
			members := m.Open("./resources/members.json")
			output := "```json\n[\n"
			for _, member := range members {
				output += printMember(member, members[len(members)-1])
			}
			output += "]\n```"
			fmt.Println(output)
			session.ChannelMessageSend(message.ChannelID, output)
		}
	}
}

func printMember(member *m.Member, lastMember *m.Member) string {
	output := "\t{"

	output += fmt.Sprintf("\"userid\": \"%s\", ", member.UserID)
	output += fmt.Sprintf("\"nick\": \"%s\", ", member.Nick)
	output += fmt.Sprintf("\"ign\": \"%s\", ", member.IGN)
	output += fmt.Sprintf("\"igid\": \"%s\", ", member.IGID)
	output += fmt.Sprintf("\"clanid\": \"%s\"", member.ClanID)

	if member == lastMember {
		output += "}\n"
	} else {
		output += "},\n"
	}

	return output
}
