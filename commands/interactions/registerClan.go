package interactions

import (
	r "calibot/commands/response"

	"github.com/bwmarrin/discordgo"
)

func RegisterClan(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {

	response := r.NewMessageResponse(AddClan(interaction).InteractionResponseData)

	return response
}
