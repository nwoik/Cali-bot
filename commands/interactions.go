package commands

import (
	"calibot/commands/interactions"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
	// Handle interaction type
	if interaction.Type == discordgo.InteractionApplicationCommand {
		switch cmd := interaction.ApplicationCommandData().Name; cmd {
		case "hello":
			response := interactions.Hello(interaction).InteractionResponse
			_ = s.InteractionRespond(interaction.Interaction, response)
		case "register":
			response := interactions.Register(interaction).InteractionResponse
			_ = s.InteractionRespond(interaction.Interaction, response)
		}
	}
}
