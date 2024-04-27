package interactions

import (
	r "calibot/commands/responses"
	e "calibot/embeds"
	"fmt"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/clan"
	m "github.com/nwoik/calibotapi/member"
)

func AddClan(clans []*c.Clan, members []*m.Member, interaction *discordgo.InteractionCreate) ([]*c.Clan, RegistrationStatus) {
	userid := interaction.Member.User.ID
	member := GetMember(members, userid)

	if member == nil {
		return clans, UserNotRegistered
	}

	args := interaction.ApplicationCommandData().Options
	name := GetArgument(args, "name").StringValue()
	clanid := GetArgument(args, "clanid").StringValue()

	if len(clanid) != 8 {
		return clans, InvalidID
	}

	clan := GetClan(clans, interaction.GuildID)
	if clan == nil {
		clan = c.CreateClan(name, clanid, interaction.GuildID).
			SetLeaderID(userid)
		clans = append(clans, clan)
		member.ClanID = clanid
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
	for _, clan := range clans {
		if clan.ClanID == id {
			return clan
		}
	}

	return nil
}

func GetClanMembers(clan *c.Clan, members []*m.Member) []*m.Member {
	clanMembers := make([]*m.Member, 0)

	for _, member := range members {
		if clan.ClanID == member.ClanID {
			clanMembers = append(clanMembers, member)
		}
	}

	return clanMembers
}

func GetMember(members []*m.Member, userid string) *m.Member {
	for _, member := range members {
		if member.UserID == userid {
			return member
		}
	}

	return nil
}

// func GetGuild(session *discordgo.Session, guildID string) *discordgo.Guild {

// }

func GetGuildMember(session *discordgo.Session, guildID string, memberID string) (*discordgo.Member, *r.Data) {
	guildMember, err := session.GuildMember(guildID, memberID)
	if err != nil {
		fmt.Println("Error retrieving member information:", err)
		return nil, r.NewResponseData("Error retrieving member information.")
	}
	return guildMember, nil
}

func MemberEmbed(member *m.Member, guildMember *discordgo.Member, discordUser *discordgo.User) *e.Embed {
	embed := e.NewRichEmbed(member.Nick, "User Info", 0x08d052c)
	embed.SetThumbnail(guildMember.AvatarURL(""))

	embed.AddField("**IGN: **", member.IGN, false)
	embed.AddField("**ID: **", member.IGID, false)

	if member.ClanID != "" {
		clans := c.Open("./resources/clan.json")
		clan := GetClan(clans, member.ClanID)
		embed.AddField("**Clan: **", clan.Name, true)
	}

	embed.SetFooter(fmt.Sprintf("Requested by %s", discordUser.Username), discordUser.AvatarURL(""))

	return embed
}

func PingUser(userid string) string {
	return fmt.Sprintf("<@%s>", userid)
}

func PingRole(id string) string {
	if len(id) < 10 {
		return ""
	}
	return fmt.Sprintf("<@&%s>", id)
}

func PrintMembers(session *discordgo.Session, clan *c.Clan, members []*m.Member, role string) string {
	var output string

	for _, member := range members {
		if isRole(session, member, clan, role) {
			output += fmt.Sprintf("%s **IGN: **%s **ID: **%s\n", PingUser(member.UserID), member.IGN, member.IGID)
		}
	}

	if output == "" {
		output = "None"
	}

	return output
}

func isRole(session *discordgo.Session, member *m.Member, clan *c.Clan, clanRole string) bool {
	guildMember, _ := GetGuildMember(session, clan.GuildID, member.UserID)

	for _, role := range guildMember.Roles {
		if role == clanRole {
			return true
		}
	}
	return false
}
