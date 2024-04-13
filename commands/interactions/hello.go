package interactions

import (
	responses "calibot/commands/responses"

	"github.com/bwmarrin/discordgo"
)

func Hello(interaction *discordgo.InteractionCreate) *responses.Response {
	data := responseData(interaction)
	response := responses.NewMessageResponse(data.InteractionResponseData)

	return response
}

func responseData(interaction *discordgo.InteractionCreate) *responses.Data {
	data := responses.NewResponseData("Hello " + interaction.Member.User.Mention())

	return data
}
