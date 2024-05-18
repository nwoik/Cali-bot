package interactions

import (
	r "calibot/commands/response"

	"github.com/bwmarrin/discordgo"
)

func RemoveWarning(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	response := r.NewMessageResponse(UnWarning(session, interaction).InteractionResponseData)

	return response
}
