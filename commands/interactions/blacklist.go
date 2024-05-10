package interactions

import (
	r "calibot/commands/response"
	"calibot/globals"

	"github.com/bwmarrin/discordgo"
	"github.com/nwoik/calibotapi/model/clan"
	m "github.com/nwoik/calibotapi/model/member"
)

func Blacklist(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	client := globals.CLIENT

	clanCollection := client.Database("calibot").Collection("clan")
	clanRepo := clan.NewClanRepo(clanCollection)
	clan, err := clanRepo.Get(interaction.GuildID)

	if err != nil {
		return r.NewMessageResponse(r.ClanNotRegisteredWithGuild().InteractionResponseData)
	}

	var data *r.Data
	clan, data = BlacklistUser(clan, session, interaction)

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

	response := r.NewMessageResponse(data.InteractionResponseData)

	return response
}
