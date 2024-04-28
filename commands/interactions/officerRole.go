package interactions

import (
	r "calibot/commands/responses"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/clan"
)

func OfficerRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clans := c.Open("./resources/clan.json")
	clan := GetClan(clans, interaction.GuildID)
	args := interaction.ApplicationCommandData().Options
	role := GetArgument(args, "role").RoleValue(session, interaction.GuildID)

	clan.OfficerRole = role.ID

	response := r.NewMessageResponse(MemberResponse().InteractionResponseData)

	c.Close("./resources/clan.json", clans)

	return response
}

func OfficerResponse() *r.Data {
	data := r.NewResponseData("Role registered")

	return data
}