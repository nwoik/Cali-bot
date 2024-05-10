package interactions

import (
	r "calibot/commands/response"

	"github.com/bwmarrin/discordgo"
)

func Promote(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clan, err := GetClan(interaction.GuildID)

	if err != nil {
		return r.NewMessageResponse(r.ClanNotRegisteredWithGuild().InteractionResponseData)
	}

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)
	member, err := GetMember(user.ID)

	if err != nil {
		return r.NewMessageResponse(r.CantPromoteNonMember().InteractionResponseData)
	}

	AddRole(session, interaction, member, clan.OfficerRole)

	response := r.NewMessageResponse(r.Promote(user).InteractionResponseData)

	return response
}
