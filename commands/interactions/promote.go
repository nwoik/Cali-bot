package interactions

import (
	r "calibot/commands/response"
	"calibot/globals"

	m "github.com/nwoik/calibotapi/model/member"

	"github.com/bwmarrin/discordgo"
)

func Promote(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	client := globals.CLIENT

	memberCollection := client.Database("calibot").Collection("member")
	memberRepo := m.NewMemberRepo(memberCollection)

	clan, err := GetClan(interaction.GuildID)

	if err != nil {
		return r.NewMessageResponse(r.ClanNotRegisteredWithGuild().InteractionResponseData)
	}

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)
	member, err := GetMember(user.ID)

	if err != nil {
		return r.NewMessageResponse(r.CantPromoteNonMember().InteractionResponseData)
	}

	AddRole(session, interaction, member, clan.OfficerRole)
	member.Rank = string(m.OFFICER)

	memberRepo.Update(member)

	response := r.NewMessageResponse(r.Promote(user).InteractionResponseData)

	return response
}
