package interactions

import (
	responses "calibot/commands/responses"

	c "github.com/nwoik/calibotapi/clan"

	"github.com/bwmarrin/discordgo"
)

func RegisterClan(session *discordgo.Session, interaction *discordgo.InteractionCreate) *responses.Response {
	clans := c.Open("./resources/clan.json")

	var status RegistrationStatus
	clans, status = AddClan(clans, interaction)

	response := responses.NewMessageResponse(ClanRegistrationResponse(interaction, status).InteractionResponseData)

	c.Close("./resources/clan.json", clans)

	return response
}

func ClanRegistrationResponse(interaction *discordgo.InteractionCreate, status RegistrationStatus) *responses.Data {
	var data *responses.Data
	args := interaction.ApplicationCommandData().Options
	name := GetArgument(args, "name").StringValue()

	switch status {
	case Success:
		data = responses.NewResponseData("Registered Clan:" + name)
	case InvalidID:
		data = responses.NewResponseData("Invalid Game-ID. Failed to register " + name)
	case AlreadyRegistered:
		data = responses.NewResponseData("Clan is already registered. Details were updated")
	case Failure:
		data = responses.NewResponseData("Failed to register. Something went wrong")
	}

	return data
}
