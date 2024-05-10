package interactions

import (
	r "calibot/commands/response"
	"calibot/globals"
	"fmt"

	"github.com/bwmarrin/discordgo"
	m "github.com/nwoik/calibotapi/model/member"
)

func RemoveMember(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	client := globals.CLIENT

	clan, err := GetClan(interaction.GuildID)

	if err != nil {
		return r.NewMessageResponse(r.NewResponseData("This server doesn't have a clan registered to it. Use `/register-clan`").InteractionResponseData)
	}

	var status Status

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)

	memberCollection := client.Database("calibot").Collection("member")
	memberRepo := m.NewMemberRepo(memberCollection)
	member, err := memberRepo.Get(user.ID)

	if err != nil {
		return r.NewMessageResponse(r.NewResponseData("This user registered.").InteractionResponseData)
	}

	member, status = RemoveClanMember(clan, member, session, interaction)
	memberRepo.Update(member)

	response := r.NewMessageResponse(RemoveMemberResponse(interaction, user, status).InteractionResponseData)

	return response
}

func RemoveMemberResponse(interaction *discordgo.InteractionCreate, user *discordgo.User, status Status) *r.Data {
	var data *r.Data

	switch status {
	case ClanMemberRemoved:
		data = r.NewResponseData(fmt.Sprintf("%s has been removed from the clan", user.Mention()))
	case ClanMemberNotFound:
		data = r.NewResponseData("This user isn't in the clan.")
	}

	return data
}
