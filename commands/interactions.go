package commands

import (
	"calibot/commands/interactions"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	// Handle interaction type
	if interaction.Type == discordgo.InteractionApplicationCommand {
		switch cmd := interaction.ApplicationCommandData().Name; cmd {
		case "register":
			response := interactions.Register(session, interaction).InteractionResponse
			_ = session.InteractionRespond(interaction.Interaction, response)

		case "registerclan":
			response := interactions.RegisterClan(session, interaction).InteractionResponse
			_ = session.InteractionRespond(interaction.Interaction, response)
		}
	}
}
