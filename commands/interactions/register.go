package interactions

import (
	responses "calibot/commands/responses"

	"github.com/bwmarrin/discordgo"
)

func Register(interaction *discordgo.InteractionCreate) *responses.Response {
	data := registrationResponseData(interaction)
	response := responses.NewMessageResponse(data.InteractionResponseData)

	return response
}

func registrationResponseData(interaction *discordgo.InteractionCreate) *responses.Data {
	data := responses.NewResponseData("Registered" + interaction.Member.User.Mention())

	return data
}
