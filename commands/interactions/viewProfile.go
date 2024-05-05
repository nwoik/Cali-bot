package interactions

import (
	r "calibot/commands/responses"

	"github.com/bwmarrin/discordgo"
	m "github.com/nwoik/calibotapi/model/member"
)

func ViewProfile(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	members := m.Open("./resources/members.json")

	args := interaction.ApplicationCommandData().Options
	var member *m.Member

	if len(args) != 0 {
		user := GetArgument(args, "member").UserValue(session)
		member = GetMember(members, user.ID)
	} else {
		member = GetMember(members, interaction.Member.User.ID)
	}

	response := r.NewMessageResponse(EmbedResponse(session, interaction, member).InteractionResponseData)

	return response
}

func EmbedResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate, member *m.Member) *r.Data {
	var data *r.Data

	if member == nil {
		data = r.NewResponseData("User is not registered with the bot.")
	} else {
		guildID := interaction.GuildID
		interactionUser := interaction.Member.User
		memberID := member.UserID

		// Get the member's information
		guildMember, errdata := GetGuildMember(session, guildID, memberID)
		if errdata != nil {
			return errdata
		}

		embed := MemberEmbed(member, guildMember, interactionUser)

		data = r.NewResponseData("").AddEmbed(embed)
	}

	return data
}
