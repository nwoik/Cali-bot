package commands

import "github.com/bwmarrin/discordgo"

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Handle interaction type
	if i.Type == discordgo.InteractionApplicationCommand {
		// Handle the "/hello" command
		if i.ApplicationCommandData().Name == "hello" {
			// Respond to the command with "hello"
			response := &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "hello",
				},
			}
			_ = s.InteractionRespond(i.Interaction, response)
		}
	}
}
