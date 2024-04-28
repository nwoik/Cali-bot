package interactions

import (
	r "calibot/commands/responses"
	"fmt"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/clan"
	m "github.com/nwoik/calibotapi/member"
)

type BlacklistStatus int

const (
	Blacklisted        BlacklistStatus = 1
	AlreadyBlacklisted BlacklistStatus = 2
)

func Blacklist(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	members := m.Open("./resources/members.json")
	clans := c.Open("./resources/clan.json")
	clan := GetClan(clans, interaction.GuildID)

	var status BlacklistStatus
	clan, status = BlacklistUser(clan, members, session, interaction)

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)
	member := GetMember(members, user.ID)

	if member != nil {
		members, _ = RemoveClanMember(clan, members, session, interaction)
	}

	response := r.NewMessageResponse(BlacklistResponse(interaction, user, status).InteractionResponseData)

	c.Close("./resources/clan.json", clans)
	m.Close("./resources/members.json", members)

	return response
}

func BlacklistResponse(interaction *discordgo.InteractionCreate, user *discordgo.User, status BlacklistStatus) *r.Data {
	var data *r.Data

	switch status {
	case Blacklisted:
		data = r.NewResponseData(fmt.Sprintf("%s has been blacklisted.", user.Mention()))
	case AlreadyBlacklisted:
		data = r.NewResponseData("This user is already blacklisted.")
	}

	return data
}
