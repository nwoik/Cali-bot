package interactions

import (
	responses "calibot/commands/responses"

	m "github.com/nwoik/calibotapi/model/member"

	"github.com/bwmarrin/discordgo"
)

func Register(session *discordgo.Session, interaction *discordgo.InteractionCreate) *responses.Response {

	members := m.Open("./resources/members.json")

	var status Status
	members, status = AddMember(members, interaction)

	// possibly for changing nicks
	// parameters := discordgo.GuildMemberParams{}
	// parameters.Nick = interaction.Member.Nick + " -> " + interaction.Member.User.ID

	// _, err := session.GuildMemberEdit(interaction.GuildID, interaction.Member.User.ID, &parameters)
	// if err != nil {
	// 	fmt.Println("Error changing member nickname:", err)
	// }

	response := responses.NewMessageResponse(RegistrationResponse(interaction, status).InteractionResponseData)

	m.Close("./resources/members.json", members)

	return response
}

func RegistrationResponse(interaction *discordgo.InteractionCreate, status Status) *responses.Data {
	var data *responses.Data

	switch status {
	case Success:
		data = responses.NewResponseData("Registered" + interaction.Member.User.Mention())
	case InvalidID:
		data = responses.NewResponseData("Invalid Game-ID. Failed to register" + interaction.Member.User.Mention())
	case AlreadyRegistered:
		data = responses.NewResponseData("User is already registered. Details were updated")
	case Failure:
		data = responses.NewResponseData("Failed to register. Something went wrong")
	}

	return data
}
