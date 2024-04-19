package commands

import (
	"github.com/bwmarrin/discordgo"
)

func RegisterCommand(s *discordgo.Session) error {

	commands := make([]*discordgo.ApplicationCommand, 0)

	register := NewChatApplicationCommand("register", "Registers a user with the bot")

	register.SetOptions(AddOption(register.GetOptions(), "ign", "Your in-game name", discordgo.ApplicationCommandOptionString, true))
	register.SetOptions(AddOption(register.GetOptions(), "gameid", "Your in-game id", discordgo.ApplicationCommandOptionString, true))

	// Add commands here
	commands = append(commands, register.ApplicationCommand)

	// Register the command globally
	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", commands)
	if err != nil {
		return err
	}

	return nil
}

func AddOption(commandOptions []*discordgo.ApplicationCommandOption, name string, description string, optionType discordgo.ApplicationCommandOptionType, required bool) []*discordgo.ApplicationCommandOption {
	commandOption := NewCommandOption(name, description, optionType, required)

	commandOptions = append(commandOptions, commandOption.ApplicationCommandOption)

	return commandOptions
}
