package events

import (
	c "calibot/client"
	"fmt"

	"github.com/bwmarrin/discordgo"
	m "github.com/nwoik/calibotapi/model/member"
)

func MemberLeave(session *discordgo.Session, guildMember *discordgo.GuildMemberRemove) {
	fmt.Printf("Member left the server: %s\n", guildMember.User.Username)

	client := c.NewMongoClient()
	collection := client.Database("calibot").Collection("members")
	memberRepo := m.NewMemberRepo(collection)

	member, err := memberRepo.Get(guildMember.User.ID)

	if err != nil {
		return
	}

	member.ClanID = ""
}
