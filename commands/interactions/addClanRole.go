package interactions

import (
	r "calibot/commands/response"

	"github.com/bwmarrin/discordgo"
)

func AddClanRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clan, err := GetClan(interaction.GuildID)

	if err != nil {
		return r.NewMessageResponse(r.ClanNotRegisteredWithGuild().InteractionResponseData)
	}

	args := interaction.ApplicationCommandData().Options
	role := GetArgument(args, "role").RoleValue(session, clan.GuildID)

	response := r.NewMessageResponse(AddExtraRole(clan, role.ID).InteractionResponseData)

	return response
}
