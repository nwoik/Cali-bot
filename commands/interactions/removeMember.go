package interactions

import (
	r "calibot/commands/response"
	"calibot/globals"

	"github.com/bwmarrin/discordgo"
	m "github.com/nwoik/calibotapi/model/member"
)

func RemoveMember(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	client := globals.CLIENT

	clan, err := GetClan(interaction.GuildID)

	if err != nil {
		return r.NewMessageResponse(r.ClanNotRegisteredWithGuild().InteractionResponseData)
	}

	var data *r.Data

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)

	memberCollection := client.Database("calibot").Collection("member")
	memberRepo := m.NewMemberRepo(memberCollection)
	member, err := memberRepo.Get(user.ID)

	if err != nil {
		return r.NewMessageResponse(r.NewResponseData("This user registered.").InteractionResponseData)
	}

	member, data = RemoveClanMember(clan, member, session, interaction)
	memberRepo.Update(member)

	response := r.NewMessageResponse(data.InteractionResponseData)

	return response
}
