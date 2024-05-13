package events

import (
	i "calibot/commands/interactions"
	"calibot/globals"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	m "github.com/nwoik/calibotapi/model/member"
)

const prefix string = "£"

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == session.State.SessionID {
		return
	}

	// channel, _ := session.Channel(message.ChannelID)
	guild, _ := session.Guild(message.GuildID)

	args := strings.Split(message.Content, " ")
	if strings.Contains(args[0], prefix) {
		fmt.Println(args)
		command := strings.Trim(args[0], prefix)

		if command == "ping" {
			session.ChannelMessageSend(message.ChannelID, "pong")
		}

		if command == "update" {
			client := globals.CLIENT

			memberCollection := client.Database("calibot").Collection("member")
			memberRepo := m.NewMemberRepo(memberCollection)

			clan, _ := i.GetClan(guild.ID)

			members, _ := i.GetMembersWithCond(i.Pred("clanid", clan.ClanID))

			for _, member := range members {
				discordUser, _ := i.GetGuildMember(session, guild.ID, member.UserID)
				member.DateJoined = discordUser.JoinedAt.Format("02/01/2006")

				memberRepo.Update(member)
			}
		}
	}
}
