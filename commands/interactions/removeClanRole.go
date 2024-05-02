package interactions

import (
	r "calibot/commands/responses"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/clan"
)

func RemoveClanRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clans := c.Open("./resources/clan.json")
	clan := GetClan(clans, interaction.GuildID)

	args := interaction.ApplicationCommandData().Options
	role := GetArgument(args, "role").RoleValue(session, clan.GuildID)

	var status Status
	clan.ExtraRoles, status = Remove(clan.ExtraRoles, role.ID)

	response := r.NewMessageResponse(RoleRemovalResponse(status).InteractionResponseData)

	c.Close("./resources/clan.json", clans)

	return response
}

func RoleRemovalResponse(status Status) *r.Data {
	var data *r.Data

	switch status {
	case Removed:
		data = r.NewResponseData("Role has been removed from clan's extra roles")
	case NotFound:
		data = r.NewResponseData("Role not found")
	}

	return data
}
