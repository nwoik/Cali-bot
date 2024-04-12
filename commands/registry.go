package commands

import (
	"github.com/bwmarrin/discordgo"
)

func RegisterCommand(s *discordgo.Session) error {

	commands := make([]*discordgo.ApplicationCommand, 0)

	// Add commands here
	commands = AddCommand(commands, "hello", "Says hello", discordgo.ChatApplicationCommand)

	// Register the command globally
	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", commands)
	if err != nil {
		return err
	}

	return nil
}

func AddCommand(commands []*discordgo.ApplicationCommand, name string, description string, commandType discordgo.ApplicationCommandType) []*discordgo.ApplicationCommand {
	commands = append(commands, NewApplicationCommand(name, description, commandType).ApplicationCommand)

	return commands
}
