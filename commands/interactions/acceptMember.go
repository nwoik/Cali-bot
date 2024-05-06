package interactions

import (
	r "calibot/commands/response"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func AcceptMember(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	status := AddClanMember(session, interaction)

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)

	response := r.NewMessageResponse(AcceptionResponse(interaction, user, status).InteractionResponseData)
	return response
}

func AcceptionResponse(interaction *discordgo.InteractionCreate, user *discordgo.User, status Status) *r.Data {
	var data *r.Data

	switch status {
	case Accepted:
		data = r.NewResponseData(fmt.Sprintf("%s has been added to the clan", user.Mention()))
	case AlreadyAccepted:
		data = r.NewResponseData("User is already in the clan.")
	case BlacklistedUser:
		data = r.NewResponseData("User is blacklisted and cannot be accepted into clan")
	case UserNotRegistered:
		data = r.NewResponseData("User is not registered with the bot. User `/register`")
	}

	return data
}
