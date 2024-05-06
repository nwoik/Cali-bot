package interactions

import (
	"calibot/client"
	r "calibot/commands/response"
	"context"

	"github.com/bwmarrin/discordgo"
	m "github.com/nwoik/calibotapi/model/member"
	"go.mongodb.org/mongo-driver/mongo"
)

func ViewProfile(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	client, err := client.NewMongoClient()

	defer client.Disconnect(context.Background())

	if err != nil {
		return r.NewMessageResponse(FaildDBResponse().InteractionResponseData)
	}

	args := interaction.ApplicationCommandData().Options
	var member *m.Member

	if len(args) != 0 {
		user := GetArgument(args, "member").UserValue(session)
		member, err = GetMember(client, user.ID)
		if err != nil {
			return r.NewMessageResponse(r.NewResponseData("This user is not registered with the bot.").InteractionResponseData)
		}
	} else {
		member, err = GetMember(client, interaction.Member.User.ID)
		if err != nil {
			return r.NewMessageResponse(r.NewResponseData("This user is not registered with the bot.").InteractionResponseData)
		}
	}

	response := r.NewMessageResponse(EmbedResponse(client, session, interaction, member).InteractionResponseData)

	return response
}

func EmbedResponse(client *mongo.Client, session *discordgo.Session, interaction *discordgo.InteractionCreate, member *m.Member) *r.Data {
	var data *r.Data

	if member == nil {
		data = r.NewResponseData("User is not registered with the bot.")
	} else {
		guildID := interaction.GuildID
		interactionUser := interaction.Member.User
		memberID := member.UserID

		// Get the member's information
		guildMember, errdata := GetGuildMember(session, guildID, memberID)
		if errdata != nil {
			return errdata
		}

		embed := MemberEmbed(client, member, guildMember, interactionUser)

		data = r.NewResponseData("").AddEmbed(embed)
	}

	return data
}
