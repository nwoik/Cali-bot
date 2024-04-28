package interactions

import (
	r "calibot/commands/responses"
	"fmt"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/clan"
	m "github.com/nwoik/calibotapi/member"
)

type AcceptionStatus int

const (
	Accepted        AcceptionStatus = 1
	AlreadyAccepted AcceptionStatus = 2
	BlacklistedUser AcceptionStatus = 3
	NotRegistered   AcceptionStatus = 4
)

func AcceptMember(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clans := c.Open("./resources/clan.json")
	members := m.Open("./resources/members.json")
	clan := GetClan(clans, interaction.GuildID)

	var status AcceptionStatus

	members, status = AddClanMember(clan, members, session, interaction)

	// possibly for changing nicks
	// parameters := discordgo.GuildMemberParams{}
	// parameters.Nick = interaction.Member.Nick + " -> " + interaction.Member.User.ID

	// _, err := session.GuildMemberEdit(interaction.GuildID, interaction.Member.User.ID, &parameters)
	// if err != nil {
	// 	fmt.Println("Error changing member nickname:", err)
	// }

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)

	response := r.NewMessageResponse(AcceptionResponse(interaction, user, status).InteractionResponseData)

	c.Close("./resources/clan.json", clans)
	m.Close("./resources/members.json", members)

	return response
}

func AcceptionResponse(interaction *discordgo.InteractionCreate, user *discordgo.User, status AcceptionStatus) *r.Data {
	var data *r.Data

	switch status {
	case Accepted:
		data = r.NewResponseData(fmt.Sprintf("%s has been added to the clan", user.Mention()))
	case AlreadyAccepted:
		data = r.NewResponseData("User is already in the clan.")
	case BlacklistedUser:
		data = r.NewResponseData("User is blacklisted and cannot be accepted into clan")
	case NotRegistered:
		data = r.NewResponseData("User is not registered with the bot. User `/register`")
	}

	return data
}
