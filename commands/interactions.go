package commands

import (
	"calibot/commands/interactions"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	var response *discordgo.InteractionResponse
	// Handle interaction type
	if interaction.Type == discordgo.InteractionApplicationCommand && interaction.GuildID != "" {
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
		case "promote":
			response = interactions.Promote(session, interaction).InteractionResponse
		case "demote":
			response = interactions.Demote(session, interaction).InteractionResponse
		case "update-profile":
			response = interactions.UpdateProfile(session, interaction).InteractionResponse
		case "warn":
			response = interactions.Warn(session, interaction).InteractionResponse
		case "remove-warning":
			response = interactions.RemoveWarning(session, interaction).InteractionResponse
		}
	} else if interaction.Type == discordgo.InteractionMessageComponent && interaction.GuildID != "" {
		switch customID := interaction.Interaction.MessageComponentData().CustomID; customID {
		case "clan_previous_button":
			response = interactions.IncPage(session, interaction, -1).InteractionResponse
		case "clan_home_button":
			response = interactions.HomePage(session, interaction).InteractionResponse
		case "clan_next_button":
			response = interactions.IncPage(session, interaction, 1).InteractionResponse
			// default:
			// 	response = &discordgo.InteractionResponse{
			// 		Data: &discordgo.InteractionResponseData{
			// 			Content: interaction.Interaction.MessageComponentData().CustomID,
			// 		},
			// 		Type: discordgo.InteractionResponseUpdateMessage,
			// 	}
		}
		response.Type = discordgo.InteractionResponseUpdateMessage

	} else {
		response = &discordgo.InteractionResponse{
			Data: &discordgo.InteractionResponseData{
				Content: "This command cannot be used in direct messages.",
			},
			Type: discordgo.InteractionResponseChannelMessageWithSource,
		}
	}
	_ = session.InteractionRespond(interaction.Interaction, response)
}
