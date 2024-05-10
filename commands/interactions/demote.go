package interactions

import (
	r "calibot/commands/response"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Demote(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clan, _ := GetClan(interaction.GuildID)

	if clan == nil {
		return r.NewMessageResponse(r.NewResponseData("This server doesn't have a clan registered to it. Use `/register-clan`").InteractionResponseData)
	}

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)
	member, _ := GetGuildMember(session, interaction.GuildID, user.ID)

	RemoveRole(session, interaction, member, clan.OfficerRole)

	response := r.NewMessageResponse(r.NewResponseData(fmt.Sprintf("%s has been demoted :cry:", user.Mention())).InteractionResponseData)

	return response
}
