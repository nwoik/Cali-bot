package interactions

import (
	r "calibot/commands/responses"
	e "calibot/embeds"
	"fmt"

	c "github.com/nwoik/calibotapi/clan"

	"github.com/bwmarrin/discordgo"
	m "github.com/nwoik/calibotapi/member"
)

func ViewClan(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	members := m.Open("./resources/members.json")
	clans := c.Open("./resources/clan.json")

	args := interaction.ApplicationCommandData().Options
	var clanid string

	if len(args) != 0 {
		clanid = GetArgument(args, "clanid").StringValue()
	} else {
		clanid = interaction.GuildID
	}

	clan := GetClan(clans, clanid)
	if clan == nil {
		return r.NewMessageResponse(r.NewResponseData("This server doesn't have a clan registered to it. Use `/register-clan`").InteractionResponseData)
	}
	clanMembers := FilterMembers(members, InClan(clan))

	response := r.NewMessageResponse(ClanEmbedResponse(session, interaction, clan, clanMembers).InteractionResponseData)

	return response
}

func ClanEmbedResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate, clan *c.Clan, members []*m.Member) *r.Data {
	var data *r.Data

	embed := e.NewRichEmbed(fmt.Sprintf("**%s (%d/50)**", clan.Name, len(members)), "", 0xffd700)

	guildID := interaction.GuildID
	guild := GetGuild(session, guildID)

	regularMembers := FilterMembers(members, And(IsMember(session, clan), Negate(IsOfficer(session, clan)), Negate(IsLeader(clan))))
	officers := FilterMembers(members, IsOfficer(session, clan))
	leader := FilterMembers(members, IsLeader(clan))

	embed.SetThumbnail(guild.IconURL(""))
	embed.AddField("", clan.ClanID, false)
	embed.AddField("**Extra Roles**", PrintExtraRoles(clan), false)
	embed.AddField("", fmt.Sprintf("**Leader: **%s", PingRole(clan.LeaderRole)), false)
	embed.AddField("", PrintMembers(leader), false)
	embed.AddField("", fmt.Sprintf("**Officers: **%s", PingRole(clan.OfficerRole)), false)
	embed.AddField("", PrintMembers(officers), false)
	embed.AddField("", fmt.Sprintf("**Members: **%s", PingRole(clan.MemberRole)), false)
	embed.AddField("", PrintMembers(regularMembers), false)
	embed.AddField("Blacklist", PrintBlacklist(clan), false)

	embed.SetFooter(fmt.Sprintf("Requested by %s", interaction.Member.User.Username), interaction.Member.User.AvatarURL(""))

	data = r.NewResponseData("").AddEmbed(embed)

	return data
}
