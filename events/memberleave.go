package events

import (
	i "calibot/commands/interactions"
	"fmt"

	m "github.com/nwoik/calibotapi/member"

	"github.com/bwmarrin/discordgo"
)

func MemberLeave(session *discordgo.Session, member *discordgo.GuildMemberRemove) {
	fmt.Printf("Member left the server: %s\n", member.User.Username)
	members := m.Open("./resources/members.json")

	botMember := i.GetMember(members, member.User.ID)
	botMember.ClanID = ""

	m.Close("./resources/members.json", members)
}
