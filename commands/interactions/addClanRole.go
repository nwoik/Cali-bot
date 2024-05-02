package interactions

import (
	r "calibot/commands/responses"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/clan"
)

func AddClanRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clans := c.Open("./resources/clan.json")
	clan := GetClan(clans, interaction.GuildID)
	if clan == nil {
		return r.NewMessageResponse(r.NewResponseData("This server doesn't have a clan registered to it. Use `/register-clan`").InteractionResponseData)
	}

	args := interaction.ApplicationCommandData().Options
	role := GetArgument(args, "role").RoleValue(session, clan.GuildID)

	var status Status
	clan.ExtraRoles, status = AddExtraRole(clan.ExtraRoles, role.ID)

	response := r.NewMessageResponse(RoleAdditionResponse(status).InteractionResponseData)

	c.Close("./resources/clan.json", clans)

	return response
}

func RoleAdditionResponse(status Status) *r.Data {
	var data *r.Data

	switch status {
	case RoleAdded:
		data = r.NewResponseData("Role has been added to clan members")
	case AlreadyAdded:
		data = r.NewResponseData("Role is already to the clan members")
	}

	return data
}
