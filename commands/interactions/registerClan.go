package interactions

import (
	r "calibot/commands/responses"

	c "github.com/nwoik/calibotapi/clan"
	m "github.com/nwoik/calibotapi/member"

	"github.com/bwmarrin/discordgo"
)

func RegisterClan(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clans := c.Open("./resources/clan.json")
	members := m.Open("./resources/members.json")

	var status Status
	clans, status = AddClan(clans, members, interaction)

	response := r.NewMessageResponse(ClanRegistrationResponse(interaction, status).InteractionResponseData)

	c.Close("./resources/clan.json", clans)
	m.Close("./resources/members.json", members)

	return response
}

func ClanRegistrationResponse(interaction *discordgo.InteractionCreate, status Status) *r.Data {
	var data *r.Data
	args := interaction.ApplicationCommandData().Options
	name := GetArgument(args, "name").StringValue()

	switch status {
	case Success:
		data = r.NewResponseData("Registered Clan: " + name)
	case InvalidID:
		data = r.NewResponseData("Invalid Game-ID. Failed to register " + name)
	case AlreadyRegistered:
		data = r.NewResponseData("Clan is already registered. Details were updated")
	case Failure:
		data = r.NewResponseData("Clan is already registered for this server.")
	case UserNotRegistered:
		data = r.NewResponseData("User must register with the bot before registering a clan. Use `/register`")
	}

	return data
}
