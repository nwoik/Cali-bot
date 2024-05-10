package interactions

import (
	r "calibot/commands/response"

	"github.com/bwmarrin/discordgo"
	m "github.com/nwoik/calibotapi/model/member"
)

func ViewProfile(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	args := interaction.ApplicationCommandData().Options
	var member *m.Member
	var err error

	if len(args) != 0 {
		user := GetArgument(args, "member").UserValue(session)
		member, err = GetMember(user.ID)
		if err != nil {
			return r.NewMessageResponse(r.NewResponseData("This user is not registered with the bot.").InteractionResponseData)
		}
	} else {
		member, err = GetMember(interaction.Member.User.ID)
		if err != nil {
			return r.NewMessageResponse(r.NewResponseData("This user is not registered with the bot.").InteractionResponseData)
		}
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
