package commands

import (
	"github.com/bwmarrin/discordgo"
)

func RegisterCommand(s *discordgo.Session) error {

	commands := make([]*discordgo.ApplicationCommand, 0)

	register := NewChatApplicationCommand("register", "Registers a user with the bot")
	register.Options = append(register.Options, NewCommandOption("ign", "Your in-game name", discordgo.ApplicationCommandOptionString, true).ApplicationCommandOption)
	register.Options = append(register.Options, NewCommandOption("gameid", "Your in-game id", discordgo.ApplicationCommandOptionString, true).ApplicationCommandOption)

	registerClan := NewChatApplicationCommand("register-clan", "Registers a clan with the bot").
		SetDefaultMemberPermissions(discordgo.PermissionManageServer)
	registerClan.Options = append(registerClan.Options, NewCommandOption("name", "Your clan's name", discordgo.ApplicationCommandOptionString, true).ApplicationCommandOption)
	registerClan.Options = append(registerClan.Options, NewCommandOption("clanid", "Your clan's in-game id", discordgo.ApplicationCommandOptionString, true).ApplicationCommandOption)

	viewMember := NewChatApplicationCommand("view-profile", "Leave member option blank to see yourself or ping the person you want to view")
	viewMember.Options = append(viewMember.Options, NewCommandOption("member", "User's @", discordgo.ApplicationCommandOptionUser, false).ApplicationCommandOption)

	viewClan := NewChatApplicationCommand("view-clan", "Leave clanid blank to see the clan for the server you're in")
	viewClan.Options = append(viewClan.Options, NewCommandOption("clanid", "In-game id of a clan", discordgo.ApplicationCommandOptionUser, false).ApplicationCommandOption)

	// Add commands here
	commands = append(commands, register.ApplicationCommand)
	commands = append(commands, registerClan.ApplicationCommand)
	commands = append(commands, viewMember.ApplicationCommand)
	commands = append(commands, viewClan.ApplicationCommand)

	// Register the command globally
	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", commands)
	if err != nil {
		return err
	}

	return nil
}
