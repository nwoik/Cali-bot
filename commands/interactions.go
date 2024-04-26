package commands

import (
	"calibot/commands/interactions"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	var response *discordgo.InteractionResponse
	// Handle interaction type
	if interaction.Type == discordgo.InteractionApplicationCommand {
		switch cmd := interaction.ApplicationCommandData().Name; cmd {
		case "register":
			response = interactions.Register(session, interaction).InteractionResponse
		case "register-clan":
			response = interactions.RegisterClan(session, interaction).InteractionResponse
		case "view-profile":
			response = interactions.ViewMember(session, interaction).InteractionResponse
		case "view-clan":
			response = interactions.ViewClan(session, interaction).InteractionResponse
		}
	}
	_ = session.InteractionRespond(interaction.Interaction, response)
}
