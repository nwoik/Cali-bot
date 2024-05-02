package commands

import (
	"github.com/bwmarrin/discordgo"
)

var globalCommands []*discordgo.ApplicationCommand = make([]*discordgo.ApplicationCommand, 0)

func RegisterCommand(s *discordgo.Session) error {
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

	memberRole := NewChatApplicationCommand("member-role", "Register the role you want your members to have").
		SetDefaultMemberPermissions(discordgo.PermissionManageServer)
	memberRole.Options = append(memberRole.Options, NewCommandOption("role", "The @ of the role", discordgo.ApplicationCommandOptionRole, true).ApplicationCommandOption)

	officerRole := NewChatApplicationCommand("officer-role", "Register the role you want your officers to have").
		SetDefaultMemberPermissions(discordgo.PermissionManageServer)
	officerRole.Options = append(officerRole.Options, NewCommandOption("role", "The @ of the role", discordgo.ApplicationCommandOptionRole, true).ApplicationCommandOption)

	leaderRole := NewChatApplicationCommand("leader-role", "Register the role for the clan leader").
		SetDefaultMemberPermissions(discordgo.PermissionManageServer)
	leaderRole.Options = append(leaderRole.Options, NewCommandOption("role", "The @ of the role", discordgo.ApplicationCommandOptionRole, true).ApplicationCommandOption)

	accept := NewChatApplicationCommand("accept-member", "Add a user to the clan").
		SetDefaultMemberPermissions(discordgo.PermissionManageRoles)
	accept.Options = append(accept.Options, NewCommandOption("user", "User's @", discordgo.ApplicationCommandOptionUser, true).ApplicationCommandOption)

	remove := NewChatApplicationCommand("remove-member", "Removes a user from the clan").
		SetDefaultMemberPermissions(discordgo.PermissionManageRoles)
	remove.Options = append(remove.Options, NewCommandOption("user", "User's @", discordgo.ApplicationCommandOptionUser, true).ApplicationCommandOption)

	blacklist := NewChatApplicationCommand("blacklist-member", "Blacklist a user from the clan").
		SetDefaultMemberPermissions(discordgo.PermissionManageRoles)
	blacklist.Options = append(blacklist.Options, NewCommandOption("user", "User's @", discordgo.ApplicationCommandOptionUser, true).ApplicationCommandOption)

	unblacklist := NewChatApplicationCommand("unblacklist-member", "Removes a user from the clan blacklist").
		SetDefaultMemberPermissions(discordgo.PermissionManageRoles)
	unblacklist.Options = append(unblacklist.Options, NewCommandOption("user", "User's @", discordgo.ApplicationCommandOptionUser, true).ApplicationCommandOption)

	addClanRole := NewChatApplicationCommand("add-clan-role", "Adds an extra role you want members to be assigned").
		SetDefaultMemberPermissions(discordgo.PermissionManageServer)
	addClanRole.Options = append(addClanRole.Options, NewCommandOption("role", "The @ of the role", discordgo.ApplicationCommandOptionRole, true).ApplicationCommandOption)

	help := NewChatApplicationCommand("help", "Lists the bot's commands")

	// Add commands here
	globalCommands = append(globalCommands, help.ApplicationCommand)
	globalCommands = append(globalCommands, register.ApplicationCommand)
	globalCommands = append(globalCommands, registerClan.ApplicationCommand)
	globalCommands = append(globalCommands, viewMember.ApplicationCommand)
	globalCommands = append(globalCommands, viewClan.ApplicationCommand)
	globalCommands = append(globalCommands, memberRole.ApplicationCommand)
	globalCommands = append(globalCommands, officerRole.ApplicationCommand)
	globalCommands = append(globalCommands, leaderRole.ApplicationCommand)
	globalCommands = append(globalCommands, accept.ApplicationCommand)
	globalCommands = append(globalCommands, remove.ApplicationCommand)
	globalCommands = append(globalCommands, blacklist.ApplicationCommand)
	globalCommands = append(globalCommands, unblacklist.ApplicationCommand)
	globalCommands = append(globalCommands, addClanRole.ApplicationCommand)

	// Register the command globally
	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", globalCommands)
	if err != nil {
		return err
	}

	return nil
}
