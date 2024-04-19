package interactions

import (
	responses "calibot/commands/responses"

	"github.com/bwmarrin/discordgo"
)

func Hello(interaction *discordgo.InteractionCreate) *responses.Response {
	options := interaction.ApplicationCommandData().Options
	data := responses.NewResponseData("Hello " + options[0].StringValue())
	response := responses.NewMessageResponse(data.InteractionResponseData)

	return response
}
