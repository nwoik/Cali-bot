package interactions

import (
	"calibot/client"
	r "calibot/commands/response"
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/nwoik/calibotapi/model/clan"
)

func MemberRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
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

	args := interaction.ApplicationCommandData().Options
	role := GetArgument(args, "role").RoleValue(session, interaction.GuildID)

	clan.MemberRole = role.ID
	clanRepo.Update(clan)

	response := r.NewMessageResponse(MemberResponse().InteractionResponseData)

	return response
}

func MemberResponse() *r.Data {
	data := r.NewResponseData("Member role registered")

	return data
}
