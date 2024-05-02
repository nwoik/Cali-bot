package interactions

import (
	r "calibot/commands/responses"
	"fmt"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/clan"
	m "github.com/nwoik/calibotapi/member"
)

func Promote(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	clans := c.Open("./resources/clan.json")
	members := m.Open("./resources/members.json")
	clan := GetClan(clans, interaction.GuildID)
	if clan == nil {
		return r.NewMessageResponse(r.NewResponseData("This server doesn't have a clan registered to it. Use `/register-clan`").InteractionResponseData)
	}

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)
	member := GetMember(members, user.ID)

	AddRole(session, interaction, member, clan.OfficerRole)

	response := r.NewMessageResponse(r.NewResponseData(fmt.Sprintf("%s has been appointed as officer :man_police_officer:", user.Mention())).InteractionResponseData)

	c.Close("./resources/clan.json", clans)
	m.Close("./resources/members.json", members)
	return response
}
