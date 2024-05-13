package interactions

import (
	r "calibot/commands/response"
	e "calibot/components/embeds"
	"fmt"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/model/clan"
	m "github.com/nwoik/calibotapi/model/member"
)

func ViewClan(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	args := interaction.ApplicationCommandData().Options
	var clanid string
	var roleInClan bool

	if len(args) != 0 {
		clanid = GetArgument(args, "clanid").StringValue()
		roleInClan = false
	} else {
		clanid = interaction.GuildID
		roleInClan = true
	}

	clan, err := GetClan(clanid)

	if err != nil {
		return r.NewMessageResponse(r.ClanNotRegisteredWithGuild().InteractionResponseData)
	}

	members, err := GetMembersWithCond(Pred("clanid", clan.ClanID))

	response := r.NewMessageResponse(ClanResponse(session, interaction, clan, members, roleInClan).InteractionResponseData)

	return response
}

func ClanEmbedResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate, clan *c.Clan, members []*m.Member, roleInClan bool) *r.Data {
	var data *r.Data

	embed := e.NewRichEmbed(fmt.Sprintf("**%s (%d/50)**", clan.Name, len(members)), "", 0xffd700)

	guildID := interaction.GuildID
	guild := GetGuild(session, guildID)

	regularMembers := FilterMembers(members, And(IsMember(session, clan), Negate(IsOfficer(session, clan)), Negate(IsLeader(clan))))
	officers := FilterMembers(members, IsOfficer(session, clan))
	leader := FilterMembers(members, IsLeader(clan))

	embed.SetThumbnail(guild.IconURL(""))
	embed.AddField("", fmt.Sprint("Clan ID: ", clan.ClanID), false)
	embed.AddField("**Extra Roles**", PrintExtraRoles(clan, roleInClan), false)
	embed.AddField("", fmt.Sprint("**Leader: 👑 **", PingRole(clan.LeaderRole)), false)
	embed.AddField("", PrintMembers(leader), false)
	embed.AddField("", fmt.Sprint("**Officers: 👮 **", PrintRole(clan.OfficerRole, roleInClan)), false)
	embed.AddField("", PrintMembers(officers), false)
	embed.AddField("", fmt.Sprint("**Members: :military_helmet: **", PrintRole(clan.MemberRole, roleInClan)), false)
	embed.AddField("", PrintMembers(regularMembers), false)
	embed.AddField("Blacklist :no_pedestrians:", PrintBlacklist(clan), false)

	embed.SetFooter(fmt.Sprintf("Requested by %s", interaction.Member.User.Username), interaction.Member.User.AvatarURL(""))

	data = r.NewResponseData("").AddEmbed(embed)

	return data
}

func ClanResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate, clan *c.Clan, members []*m.Member, roleInClan bool) *r.Data {
	var data *r.Data

	output := ""
	output += fmt.Sprintf("**%s (%d/50)**\n", clan.Name, len(members))

	// guildID := interaction.GuildID
	// guild := GetGuild(session, guildID)

	// regularMembers := FilterMembers(members, And(IsMember(session, clan), Negate(IsOfficer(session, clan)), Negate(IsLeader(clan))))
	// officers := FilterMembers(members, IsOfficer(session, clan))
	// leader := FilterMembers(members, IsLeader(clan))

	output += fmt.Sprintf("Clan ID: %s\n", clan.ClanID)
	output += fmt.Sprintf("**Extra Roles: **%s\n", PrintExtraRoles(clan, roleInClan))

	// output += fmt.Sprintf("**Leader: 👑 **%s\n", PrintRole(clan.LeaderRole, roleInClan))
	// output += fmt.Sprintf(PrintMembers(leader))
	// output += "\n"

	// output += fmt.Sprintf("**Officers: 👮 **%s\n", PrintRole(clan.OfficerRole, roleInClan))
	// output += fmt.Sprintf(PrintMembers(officers))
	// output += "\n"

	output += fmt.Sprintf("**Members: :military_helmet: **%s\n", PrintRole(clan.MemberRole, roleInClan))
	output += fmt.Sprintf(PrintMembers(members))
	output += "\n"

	output += fmt.Sprint("**Blacklist :no_pedestrians: **\n")
	output += PrintBlacklist(clan)

	data = r.NewResponseData(output)

	return data
}
