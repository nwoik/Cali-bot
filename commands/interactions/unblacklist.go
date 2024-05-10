package interactions

import (
	r "calibot/commands/response"
	"calibot/globals"

	"github.com/bwmarrin/discordgo"
	"github.com/nwoik/calibotapi/model/clan"
)

func Unblacklist(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	client := globals.CLIENT

	clanCollection := client.Database("calibot").Collection("clan")
	clanRepo := clan.NewClanRepo(clanCollection)
	clan, err := clanRepo.Get(interaction.GuildID)

	if err != nil {
		return r.NewMessageResponse(r.ClanNotRegisteredWithGuild().InteractionResponseData)
	}

	args := interaction.ApplicationCommandData().Options
	user := GetArgument(args, "user").UserValue(session)

	var response *r.Response

	if IsBlacklisted(clan, user.ID) {
		clan.Blacklist, _ = Remove(clan.Blacklist, user.ID)
		clanRepo.Update(clan)
		response = r.NewMessageResponse(r.RemovedFromBlackList().InteractionResponseData)
	} else {
		response = r.NewMessageResponse(r.MemberNotBlacklisted().InteractionResponseData)
	}

	return response
}
