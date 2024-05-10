package interactions

import (
	r "calibot/commands/response"
	e "calibot/embeds"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Help(session *discordgo.Session, interaction *discordgo.InteractionCreate, commands []*discordgo.ApplicationCommand) *r.Response {
	data := HelpResponse(session, interaction, commands)
	response := r.NewMessageResponse(data.InteractionResponseData)

	return response
}

func HelpResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate, commands []*discordgo.ApplicationCommand) *r.Data {
	var data *r.Data
	embed := e.NewRichEmbed("**Commands**", "All the info on the bot's commands ", 0xff00e4)

	for _, command := range commands {
		if hasPermission(interaction.Member, command.DefaultMemberPermissions) {
			embed.AddField(fmt.Sprintf("**/%s**", command.Name), command.Description, false)
		}
	}

	data = r.NewResponseData("").AddEmbed(embed)

	return data
}

func hasPermission(member *discordgo.Member, requiredPermission *int64) bool {
	// Check if member or requiredPermission is nil
	if member == nil || requiredPermission == nil {
		return false
	}

	// Check if the member has the required permission
	return member.Permissions&(*requiredPermission) == *requiredPermission
}
