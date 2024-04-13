package commands

import (
	"calibot/commands/interactions"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
	// Handle interaction type
	if interaction.Type == discordgo.InteractionApplicationCommand {
		// Handle the "/hello" command
		if interaction.ApplicationCommandData().Name == "hello" {
			// Respond to the command with "hello"
			response := interactions.Hello(interaction).InteractionResponse
			_ = s.InteractionRespond(interaction.Interaction, response)
		}
	}
}
