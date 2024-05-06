package interactions

import (
	r "calibot/commands/responses"

	"github.com/bwmarrin/discordgo"
)

func RegisterClan(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {

	status := AddClan(interaction)

	response := r.NewMessageResponse(ClanRegistrationResponse(interaction, status).InteractionResponseData)

	return response
}

func ClanRegistrationResponse(interaction *discordgo.InteractionCreate, status Status) *r.Data {
	var data *r.Data
	args := interaction.ApplicationCommandData().Options
	name := GetArgument(args, "name").StringValue()

	switch status {
	case FailedDBConnection:
		data = r.NewResponseData("Failed to connect to database. Please ping admins")
	case Success:
		data = r.NewResponseData("Registered Clan: " + name + "\nUse `/viewclan` to see details\nMake sure to use `/leaderrole`, `/officerrole` and `/memberrole` for the roles you want members to have.")
	case InvalidID:
		data = r.NewResponseData("Invalid Game-ID. Failed to register " + name)
	case ClanAlreadyRegistered:
		data = r.NewResponseData("Clan is already registered. Details were updated")
	case UserNotRegistered:
		data = r.NewResponseData("User must register with the bot before registering a clan. Use `/register`")
	}

	return data
}
