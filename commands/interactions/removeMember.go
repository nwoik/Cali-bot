package interactions

import (
	r "calibot/commands/responses"
	"fmt"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/clan"
	m "github.com/nwoik/calibotapi/member"
)

func RemoveMember(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clans := c.Open("./resources/clan.json")
	members := m.Open("./resources/members.json")
	clan := GetClan(clans, interaction.GuildID)
	if clan == nil {
		return r.NewMessageResponse(r.NewResponseData("This server doesn't have a clan registered to it. Use `/register-clan`").InteractionResponseData)
	}

	var status Status

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

func RemoveMemberResponse(interaction *discordgo.InteractionCreate, user *discordgo.User, status Status) *r.Data {
	var data *r.Data

	switch status {
	case Removed:
		data = r.NewResponseData(fmt.Sprintf("%s has been removed from the clan", user.Mention()))
	case NotFound:
		data = r.NewResponseData("This user isn't in the clan.")
	}

	return data
}
