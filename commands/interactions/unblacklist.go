package interactions

import (
	r "calibot/commands/responses"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/clan"
)

func Unblacklist(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clans := c.Open("./resources/clan.json")
	clan := GetClan(clans, interaction.GuildID)

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)

	var response *r.Response

	if isBlacklisted(clan, user.ID) {
		clan.Blacklist = Remove(clan.Blacklist, user.ID)
		response = r.NewMessageResponse(r.NewResponseData("User has been removed from clan blacklist").InteractionResponseData)
	} else {
		response = r.NewMessageResponse(r.NewResponseData("User is not blacklisted").InteractionResponseData)
	}

	c.Close("./resources/clan.json", clans)
	return response
}
