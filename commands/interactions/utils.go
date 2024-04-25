package interactions

import (
	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/clan"
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

func GetClan(clans []*c.Clan, id string) *c.Clan {
	for _, clan := range clans {
		if clan.GuildID == id {
			return clan
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

func AddClan(clans []*c.Clan, interaction *discordgo.InteractionCreate) ([]*c.Clan, RegistrationStatus) {
	args := interaction.ApplicationCommandData().Options
	name := GetArgument(args, "name").StringValue()
	clanid := GetArgument(args, "clanid").StringValue()

	if len(clanid) != 8 {
		return clans, InvalidID
	}

	clan := GetClan(clans, interaction.GuildID)
	if clan == nil {
		clan = c.CreateClan(name, clanid, interaction.GuildID)
		clans = append(clans, clan)
		return clans, Success
	}
	return clans, Failure
}

func AddMember(members []*m.Member, interaction *discordgo.InteractionCreate) ([]*m.Member, RegistrationStatus) {
	args := interaction.ApplicationCommandData().Options
	gameid := GetArgument(args, "gameid").StringValue()
	ign := GetArgument(args, "ign").StringValue()

	if len(gameid) != 9 {
		return members, InvalidID
	}

	member := GetMember(members, interaction.Member.User.ID)

	if member == nil {
		member = m.CreateMember(interaction.Member.User.Username, ign, gameid, interaction.Member.User.ID)
		members = append(members, member)
		return members, Success
	}

	member.IGN = ign
	member.IGID = gameid
	return members, AlreadyRegistered

	// return members, Failure
}
