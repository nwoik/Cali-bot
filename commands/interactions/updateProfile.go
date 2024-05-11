package interactions

import (
	responses "calibot/commands/response"

	"github.com/bwmarrin/discordgo"
)

func UpdateProfile(session *discordgo.Session, interaction *discordgo.InteractionCreate) *responses.Response {
	// If we want to change user nicks when they register, here's the place

	response := responses.NewMessageResponse(UpdateMember(interaction).InteractionResponseData)

	return response
}
