package interactions

import (
	r "calibot/commands/response"
	"calibot/globals"

	"github.com/bwmarrin/discordgo"
	"github.com/nwoik/calibotapi/model/clan"
)

func OfficerRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	client := globals.CLIENT

	clanCollection := client.Database("calibot").Collection("clan")
	clanRepo := clan.NewClanRepo(clanCollection)
	clan, err := clanRepo.Get(interaction.GuildID)

	if err != nil {
		return r.NewMessageResponse(r.ClanNotRegisteredWithGuild().InteractionResponseData)
	}

	args := interaction.ApplicationCommandData().Options
	role := GetArgument(args, "role").RoleValue(session, interaction.GuildID)

	clan.OfficerRole = role.ID
	clanRepo.Update(clan)

	response := r.NewMessageResponse(r.OfficerRoleResponse().InteractionResponseData)

	return response
}
