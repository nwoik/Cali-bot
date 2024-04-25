package interactions

import (
	"github.com/bwmarrin/discordgo"
	m "github.com/nwoik/calibotapi/member"
)

func GetArgument(options []*discordgo.ApplicationCommandInteractionDataOption, name string) *discordgo.ApplicationCommandInteractionDataOption {
	for _, option := range options {
		if option.Name == name {
			return option
		}
	}
	return nil
}

func GetMember(members []*m.Member, userid string) *m.Member {
	for _, member := range members {
		if member.UserID == userid {
			return member
		}
	}

	return nil
}

func AddMember(members []*m.Member, interaction *discordgo.InteractionCreate, session *discordgo.Session) ([]*m.Member, RegistrationStatus) {
	args := interaction.ApplicationCommandData().Options
	gameid := GetArgument(args, "gameid")
	ign := GetArgument(args, "ign")

	if len(gameid.StringValue()) != 9 {
		return members, InvalidID
	}

	member := GetMember(members, interaction.Member.User.ID)

	if member == nil {
		// parameters := discordgo.GuildMemberParams{}
		// parameters.Nick = interaction.Member.Nick + " -> " + interaction.Member.User.ID

		// member, err := session.GuildMemberEdit(interaction.GuildID, interaction.Member.User.ID, &parameters)
		// if err != nil {
		// 	fmt.Println("Error changing member nickname:", err)
		// 	return members, Failure
		// }
		members = append(members, m.CreateMember(interaction.Member.User.Username, ign.StringValue(), gameid.StringValue(), interaction.Member.User.ID))
		return members, Success
	} else if member != nil {
		member.IGN = ign.StringValue()
		member.IGID = gameid.StringValue()
		return members, AlreadyRegistered
	}
	return members, Failure
}
