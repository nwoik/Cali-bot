package commands

import (
	"github.com/bwmarrin/discordgo"
)

func RegisterCommand(s *discordgo.Session) error {

	commands := make([]*discordgo.ApplicationCommand, 0)

	// Add commands here
	hello := NewChatApplicationCommand("hello", "Says hello")

	register := NewChatApplicationCommand("register", "Registers a user with the bot")
	options := []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "name",
			Description: "Your name",
			Required:    true,
		},
	}

	hello.SetOptions(options)
	// register.AddOption("hello", "Your in-game name", discordgo.ApplicationCommandOptionString, true)
	// register.AddOption("Game-ID", "Your in-game ID", discordgo.ApplicationCommandOptionInteger, true)

	commands = append(commands, hello.ApplicationCommand)
	commands = append(commands, register.ApplicationCommand)

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
