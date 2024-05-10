package interactions

import (
	r "calibot/commands/response"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Promote(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clan, err := GetClan(interaction.GuildID)

	if err != nil {
		return r.NewMessageResponse(r.NewResponseData("This server doesn't have a clan registered to it. Use `/register-clan`").InteractionResponseData)
	}

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)
	member, err := GetMember(user.ID)

	if err != nil {
		return r.NewMessageResponse(r.NewResponseData("This user is not registered with the bot.\nThey must register with the bot and clan to be an officer").InteractionResponseData)
	}

	AddRole(session, interaction, member, clan.OfficerRole)

	response := r.NewMessageResponse(r.NewResponseData(fmt.Sprintf("%s has been appointed as officer :man_police_officer:", user.Mention())).InteractionResponseData)

	return response
}
