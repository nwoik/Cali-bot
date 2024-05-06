package interactions

import (
	responses "calibot/commands/responses"

	"github.com/bwmarrin/discordgo"
)

func Register(session *discordgo.Session, interaction *discordgo.InteractionCreate) *responses.Response {
	status := AddMember(interaction)

	// If we want to change user nicks when they register, here's the place

	response := responses.NewMessageResponse(RegistrationResponse(interaction, status).InteractionResponseData)

	return response
}

func RegistrationResponse(interaction *discordgo.InteractionCreate, status Status) *responses.Data {
	var data *responses.Data

	switch status {
	case Success:
		data = responses.NewResponseData("Registered" + interaction.Member.User.Mention())
	case InvalidID:
		data = responses.NewResponseData("Invalid Game-ID. Failed to register" + interaction.Member.User.Mention())
	case UserAlreadyRegistered:
		data = responses.NewResponseData("User is already registered. Details were updated")
	}

	return data
}
