package commands

import (
	"calibot/commands/interactions"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	var response *discordgo.InteractionResponse
	// Handle interaction type
	if interaction.Type == discordgo.InteractionApplicationCommand {
		switch cmd := interaction.ApplicationCommandData().Name; cmd {
		case "help":
			response = interactions.Help(session, interaction, globalCommands).InteractionResponse
		case "register":
			response = interactions.Register(session, interaction).InteractionResponse
		case "register-clan":
			response = interactions.RegisterClan(session, interaction).InteractionResponse
		case "view-profile":
			response = interactions.ViewProfile(session, interaction).InteractionResponse
		case "view-clan":
			response = interactions.ViewClan(session, interaction).InteractionResponse
		case "member-role":
			response = interactions.MemberRole(session, interaction).InteractionResponse
		case "officer-role":
			response = interactions.OfficerRole(session, interaction).InteractionResponse
		case "leader-role":
			response = interactions.LeaderRole(session, interaction).InteractionResponse
		case "accept-member":
			response = interactions.AcceptMember(session, interaction).InteractionResponse
		case "remove-member":
			response = interactions.RemoveMember(session, interaction).InteractionResponse
		case "blacklist-member":
			response = interactions.Blacklist(session, interaction).InteractionResponse
		case "unblacklist-member":
			response = interactions.Unblacklist(session, interaction).InteractionResponse
		case "add-clan-role":
			response = interactions.AddClanRole(session, interaction).InteractionResponse
		case "remove-clan-role":
			response = interactions.RemoveClanRole(session, interaction).InteractionResponse
		case "appoint":
			response = interactions.Appoint(session, interaction).InteractionResponse
		}
	}
	_ = session.InteractionRespond(interaction.Interaction, response)
}
