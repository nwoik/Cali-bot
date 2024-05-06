package interactions

import (
	"calibot/client"
	r "calibot/commands/response"
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/nwoik/calibotapi/model/clan"
)

func Unblacklist(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
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
	user := GetArgument(args, "user").UserValue(session)

	var response *r.Response

	if IsBlacklisted(clan, user.ID) {
		clan.Blacklist, _ = Remove(clan.Blacklist, user.ID)
		clanRepo.Update(clan)
		response = r.NewMessageResponse(r.NewResponseData("User has been removed from clan blacklist").InteractionResponseData)
	} else {
		response = r.NewMessageResponse(r.NewResponseData("User is not blacklisted").InteractionResponseData)
	}

	return response
}
