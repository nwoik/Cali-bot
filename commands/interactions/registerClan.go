package interactions

import (
	responses "calibot/commands/responses"

	c "github.com/nwoik/calibotapi/clan"
	m "github.com/nwoik/calibotapi/member"

	"github.com/bwmarrin/discordgo"
)

func RegisterClan(session *discordgo.Session, interaction *discordgo.InteractionCreate) *responses.Response {
	clans := c.Open("./resources/clan.json")
	members := m.Open("./resources/members.json")

	var status RegistrationStatus
	clans, status = AddClan(clans, members, interaction)

	response := responses.NewMessageResponse(ClanRegistrationResponse(interaction, status).InteractionResponseData)

	c.Close("./resources/clan.json", clans)
	m.Close("./resources/members.json", members)

	return response
}

func ClanRegistrationResponse(interaction *discordgo.InteractionCreate, status RegistrationStatus) *responses.Data {
	var data *responses.Data
	args := interaction.ApplicationCommandData().Options
	name := GetArgument(args, "name").StringValue()

	switch status {
	case Success:
		data = responses.NewResponseData("Registered Clan: " + name)
	case InvalidID:
		data = responses.NewResponseData("Invalid Game-ID. Failed to register " + name)
	case AlreadyRegistered:
		data = responses.NewResponseData("Clan is already registered. Details were updated")
	case Failure:
		data = responses.NewResponseData("Clan is already registered for this server.")
	case UserNotRegistered:
		data = responses.NewResponseData("User must register with the bot before registering a clan. Use `/register`")
	}

	return data
}
