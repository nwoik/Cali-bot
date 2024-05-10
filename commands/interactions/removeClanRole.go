package interactions

import (
	r "calibot/commands/response"
	"calibot/globals"

	"github.com/bwmarrin/discordgo"
	"github.com/nwoik/calibotapi/model/clan"
)

func RemoveClanRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	client := globals.CLIENT

	clanCollection := client.Database("calibot").Collection("clan")
	clanRepo := clan.NewClanRepo(clanCollection)
	clan, err := clanRepo.Get(interaction.GuildID)

	if err != nil {
		return r.NewMessageResponse(r.ClanNotRegisteredWithGuild().InteractionResponseData)
	}

	args := interaction.ApplicationCommandData().Options
	role := GetArgument(args, "role").RoleValue(session, clan.GuildID)

	var status Status
	clan.ExtraRoles, status = Remove(clan.ExtraRoles, role.ID)
	clanRepo.Update(clan)

	response := r.NewMessageResponse(RoleRemovalResponse(status).InteractionResponseData)

	return response
}

func RoleRemovalResponse(status Status) *r.Data {
	var data *r.Data

	switch status {
	case Removed:
		data = r.RoleRemoved()
	case NotFound:
		data = r.RoleNotFound()
	}

	return data
}
