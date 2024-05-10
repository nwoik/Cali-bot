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
		return r.NewMessageResponse(r.NewResponseData("This server doesn't have a clan registered to it. Use `/register-clan`").InteractionResponseData)
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
	case ClanMemberRemoved:
		data = r.NewResponseData("Role has been removed from clan's extra roles")
	case ClanMemberNotFound:
		data = r.NewResponseData("Role not found")
	}

	return data
}
