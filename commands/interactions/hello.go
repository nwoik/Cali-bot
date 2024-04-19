package interactions

import (
	responses "calibot/commands/responses"

	"github.com/bwmarrin/discordgo"
)

func Hello(interaction *discordgo.InteractionCreate) *responses.Response {
	data := responses.NewResponseData("Hello " + interaction.Member.User.Mention())
	response := responses.NewMessageResponse(data.InteractionResponseData)

	return response
}
