package interactions

import (
	r "calibot/commands/responses"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/model/clan"
)

func Unblacklist(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clans := c.Open("./resources/clan.json")
	clan := GetClan(clans, interaction.GuildID)
	if clan == nil {
		return r.NewMessageResponse(r.NewResponseData("This server doesn't have a clan registered to it. Use `/register-clan`").InteractionResponseData)
	}

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)

	var response *r.Response

	if IsBlacklisted(clan, user.ID) {
		clan.Blacklist, _ = Remove(clan.Blacklist, user.ID)
		response = r.NewMessageResponse(r.NewResponseData("User has been removed from clan blacklist").InteractionResponseData)
	} else {
		response = r.NewMessageResponse(r.NewResponseData("User is not blacklisted").InteractionResponseData)
	}

	c.Close("./resources/clan.json", clans)
	return response
}
