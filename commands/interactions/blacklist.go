package interactions

import (
	"calibot/client"
	r "calibot/commands/responses"
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nwoik/calibotapi/model/clan"
	m "github.com/nwoik/calibotapi/model/member"
)

func Blacklist(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	client, err := client.NewMongoClient()

	defer client.Disconnect(context.Background())

	if err != nil {
		return r.NewMessageResponse(FaildDBResponse().InteractionResponseData)
	}

	clanCollection := client.Database("calibot").Collection("clan")
	clanRepo := clan.NewClanRepo(clanCollection)
	clan, err := clanRepo.Get(interaction.GuildID)

	if err != nil {
		return r.NewMessageResponse(r.NewResponseData("This server doesn't have a clan registered to it. Use `/register-clan`").InteractionResponseData)
	}

	var status Status
	clan, status = BlacklistUser(clan, session, interaction)

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)

	memberCollection := client.Database("calibot").Collection("member")
	memberRepo := m.NewMemberRepo(memberCollection)
	member, err := memberRepo.Get(user.ID)

	if member != nil {
		member, _ = RemoveClanMember(clan, member, session, interaction)
		memberRepo.Update(member)
	}

	clanRepo.Update(clan)

	response := r.NewMessageResponse(BlacklistResponse(interaction, user, status).InteractionResponseData)

	return response
}

func BlacklistResponse(interaction *discordgo.InteractionCreate, user *discordgo.User, status Status) *r.Data {
	var data *r.Data

	switch status {
	case Blacklisted:
		data = r.NewResponseData(fmt.Sprintf("%s has been blacklisted.", user.Mention()))
	case AlreadyBlacklisted:
		data = r.NewResponseData("This user is already blacklisted.")
	}

	return data
}

func FaildDBResponse() *r.Data {
	return r.NewResponseData("Failed to connect to database")
}
