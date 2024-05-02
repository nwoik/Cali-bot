package interactions

import (
	r "calibot/commands/responses"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/clan"
)

func LeaderRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clans := c.Open("./resources/clan.json")
	clan := GetClan(clans, interaction.GuildID)

	if clan == nil {
		return r.NewMessageResponse(r.NewResponseData("This server doesn't have a clan registered to it. Use `/register-clan`").InteractionResponseData)
	}
	args := interaction.ApplicationCommandData().Options
	role := GetArgument(args, "role").RoleValue(session, interaction.GuildID)

	clan.LeaderRole = role.ID

	response := r.NewMessageResponse(MemberResponse().InteractionResponseData)

	c.Close("./resources/clan.json", clans)

	return response
}

func LeaderResponse() *r.Data {
	data := r.NewResponseData("Role registered")

	return data
}
