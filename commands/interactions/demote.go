package interactions

import (
	r "calibot/commands/response"

	"github.com/bwmarrin/discordgo"
)

func Demote(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clan, _ := GetClan(interaction.GuildID)

	if clan == nil {
		return r.NewMessageResponse(r.ClanNotRegistered().InteractionResponseData)
	}

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)
	member, _ := GetGuildMember(session, interaction.GuildID, user.ID)

	RemoveRole(session, interaction, member, clan.OfficerRole)

	response := r.NewMessageResponse(r.OfficerDemoted(user).InteractionResponseData)

	return response
}
