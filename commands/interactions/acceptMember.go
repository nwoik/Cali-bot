package interactions

import (
	r "calibot/commands/response"

	"github.com/bwmarrin/discordgo"
)

func AcceptMember(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {

	response := r.NewMessageResponse(AddClanMember(session, interaction).InteractionResponseData)
	return response
}
