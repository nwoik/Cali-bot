package interactions

import (
	r "calibot/commands/responses"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/clan"
)

func MemberRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clans := c.Open("./resources/clan.json")
	clan := GetClan(clans, interaction.GuildID)
	args := interaction.ApplicationCommandData().Options
	role := GetArgument(args, "role").RoleValue(session, interaction.GuildID)

	clan.MemberRole = role.ID

	response := r.NewMessageResponse(MemberResponse().InteractionResponseData)

	c.Close("./resources/clan.json", clans)

	return response
}

func MemberResponse() *r.Data {
	data := r.NewResponseData("Role registered")

	return data
}
