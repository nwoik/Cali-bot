package interactions

import (
	r "calibot/commands/response"

	"github.com/bwmarrin/discordgo"
)

func AddClanRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clan, err := GetClan(interaction.GuildID)

	if err != nil {
		return r.NewMessageResponse(r.NewResponseData("This server doesn't have a clan registered to it. Use `/register-clan`").InteractionResponseData)
	}

	args := interaction.ApplicationCommandData().Options
	role := GetArgument(args, "role").RoleValue(session, clan.GuildID)

	status := AddExtraRole(clan, role.ID)

	response := r.NewMessageResponse(RoleAdditionResponse(status).InteractionResponseData)

	return response
}

func RoleAdditionResponse(status Status) *r.Data {
	var data *r.Data

	switch status {
	case RoleAdded:
		data = r.NewResponseData("Role has been added to clan members")
	case AlreadyAdded:
		data = r.NewResponseData("Role is already to the clan members")
	case ClanNotRegistered:
		return r.NewResponseData("This server doesn't have a clan registered to it. Use `/register-clan`")
	}

	return data
}
