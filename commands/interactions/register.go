package interactions

import (
	responses "calibot/commands/response"

	"github.com/bwmarrin/discordgo"
)

func Register(session *discordgo.Session, interaction *discordgo.InteractionCreate) *responses.Response {
	// If we want to change user nicks when they register, here's the place

	response := responses.NewMessageResponse(AddMember(interaction).InteractionResponseData)

	return response
}
