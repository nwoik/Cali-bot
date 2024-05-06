package interactions

import (
	"calibot/client"
	r "calibot/commands/responses"
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/nwoik/calibotapi/model/clan"
)

func LeaderRole(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
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

	clan.LeaderRole = role.ID
	clanRepo.Update(clan)

	response := r.NewMessageResponse(LeaderResponse().InteractionResponseData)

	return response
}

func LeaderResponse() *r.Data {
	data := r.NewResponseData("Leader role registered")

	return data
}
