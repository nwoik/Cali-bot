package interactions

import (
	responses "calibot/commands/responses"

	clans "github.com/nwoik/calibotapi/clans"

	"github.com/bwmarrin/discordgo"
)

func Register(interaction *discordgo.InteractionCreate) *responses.Response {

	members := clans.NewMembers().Open("./resources/members.json")

	members = members.SetMembers(AddMember(members.GetMembers(), interaction))

	members.Close("./resources/members.json")

	response := responses.NewMessageResponse(successfulRegistrationResponse(interaction).InteractionResponseData)

	return response
}

func successfulRegistrationResponse(interaction *discordgo.InteractionCreate) *responses.Data {
	data := responses.NewResponseData("Registered" + interaction.Member.User.Mention())

	return data
}
