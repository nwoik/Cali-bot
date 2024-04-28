package interactions

import (
	r "calibot/commands/responses"
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
		embed.AddField(fmt.Sprintf("**/%s**", command.Name), command.Description, false)
	}

	data = r.NewResponseData("").AddEmbed(embed)

	return data
}
