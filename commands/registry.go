package commands

import (
	"github.com/bwmarrin/discordgo"
)

func RegisterCommand(s *discordgo.Session) error {

	commands := make([]*discordgo.ApplicationCommand, 0)

	register := NewChatApplicationCommand("register", "Registers a user with the bot")
	register.Options = append(register.Options, NewCommandOption("ign", "Your in-game name", discordgo.ApplicationCommandOptionString, true).ApplicationCommandOption)
	register.Options = append(register.Options, NewCommandOption("gameid", "Your in-game id", discordgo.ApplicationCommandOptionString, true).ApplicationCommandOption)

	registerClan := NewChatApplicationCommand("registerclan", "Registers a clan with the bot")
	registerClan.Options = append(registerClan.Options, NewCommandOption("name", "Your clan's name", discordgo.ApplicationCommandOptionString, true).ApplicationCommandOption)
	registerClan.Options = append(registerClan.Options, NewCommandOption("clanid", "Your clan's in-game id", discordgo.ApplicationCommandOptionString, true).ApplicationCommandOption)

	// Add commands here
	commands = append(commands, register.ApplicationCommand)
	commands = append(commands, registerClan.ApplicationCommand)

	// Register the command globally
	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", commands)
	if err != nil {
		return err
	}

	return nil
}
