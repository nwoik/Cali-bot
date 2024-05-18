package interactions

import (
	r "calibot/commands/response"

	"github.com/bwmarrin/discordgo"
)

func Warn(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	response := r.NewMessageResponse(WarnUser(session, interaction).InteractionResponseData)

	return response
}
