package interactions

import (
	r "calibot/commands/responses"
	"fmt"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/clan"
	m "github.com/nwoik/calibotapi/member"
)

type RemovalStatus int

const (
	Removed          RemovalStatus = 1
	MemberNotPresent RemovalStatus = 2
)

func RemoveMember(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clans := c.Open("./resources/clan.json")
	members := m.Open("./resources/members.json")
	clan := GetClan(clans, interaction.GuildID)

	var status RemovalStatus

	members, status = RemoveClanMember(clan, members, session, interaction)

	// possibly for changing nicks
	// parameters := discordgo.GuildMemberParams{}
	// parameters.Nick = interaction.Member.Nick + " -> " + interaction.Member.User.ID

	// _, err := session.GuildMemberEdit(interaction.GuildID, interaction.Member.User.ID, &parameters)
	// if err != nil {
	// 	fmt.Println("Error changing member nickname:", err)
	// }

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)

	response := r.NewMessageResponse(RemoveMemberResponse(interaction, user, status).InteractionResponseData)

	c.Close("./resources/clan.json", clans)
	m.Close("./resources/members.json", members)

	return response
}

func RemoveMemberResponse(interaction *discordgo.InteractionCreate, user *discordgo.User, status RemovalStatus) *r.Data {
	var data *r.Data

	switch status {
	case Removed:
		data = r.NewResponseData(fmt.Sprintf("%s has been removed from the clan", user.Mention()))
	case MemberNotPresent:
		data = r.NewResponseData("This user isn't in the clan.")
	}

	return data
}
