package commands

import (
	"calibot/commands/interactions"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Handle interaction type
	if i.Type == discordgo.InteractionApplicationCommand {
		// Handle the "/hello" command
		if i.ApplicationCommandData().Name == "hello" {
			// Respond to the command with "hello"
			response := interactions.Hello().InteractionResponse
			_ = s.InteractionRespond(i.Interaction, response)
		}
	}
}
