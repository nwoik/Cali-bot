package interactions

import (
	"github.com/bwmarrin/discordgo"
	clans "github.com/nwoik/calibotapi/clans"
	cUtils "github.com/nwoik/calibotapi/utils"
)

func GetArgument(options []*discordgo.ApplicationCommandInteractionDataOption, name string) *discordgo.ApplicationCommandInteractionDataOption {
	for _, option := range options {
		if option.Name == name {
			return option
		}
	}
	return nil
}

func GetMember(members []*clans.Member, userid string) *clans.Member {
	for _, member := range members {
		if member.UserID == userid {
			return member
		}
	}

	return nil
}

func AddMember(members []*clans.Member, interaction *discordgo.InteractionCreate) []*clans.Member {

	member := GetMember(members, interaction.Member.User.ID)

	if member == nil {
		args := interaction.ApplicationCommandData().Options
		ign := GetArgument(args, "ign")
		gameid := GetArgument(args, "gameid")

		members = cUtils.AddMember(members, interaction.Member.Nick, ign.StringValue(), gameid.StringValue(), interaction.Member.User.ID)
	}

	return members
}
