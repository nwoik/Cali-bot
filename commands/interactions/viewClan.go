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
	clanMembers := GetClanMembers(clan, members)

	response := r.NewMessageResponse(ClanEmbedResponse(session, interaction, clan, clanMembers).InteractionResponseData)

	return response
}

func ClanEmbedResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate, clan *c.Clan, members []*m.Member) *r.Data {
	var data *r.Data
	embed := e.NewRichEmbed(fmt.Sprintf("**%s (%d/50)**", clan.Name, len(GetClanMembers(clan, members))), "", 0xffd700)

	guildID := interaction.GuildID
	guild := GetGuild(session, guildID)

	embed.SetThumbnail(guild.IconURL(""))
	embed.AddField("**Extra Roles**", PrintExtraRoles(clan), false)
	embed.AddField("", fmt.Sprintf("**Leader: **%s", PingRole(clan.LeaderRole)), false)
	embed.AddField("", PrintMembers(session, clan, members, clan.LeaderRole), false)
	embed.AddField("", fmt.Sprintf("**Officers: **%s", PingRole(clan.OfficerRole)), false)
	embed.AddField("", PrintMembers(session, clan, members, clan.OfficerRole), false)
	embed.AddField("", fmt.Sprintf("**Members: **%s", PingRole(clan.MemberRole)), false)
	embed.AddField("", PrintMembers(session, clan, members, clan.MemberRole), false)
	embed.AddField("Blacklist", PrintBlacklist(clan), false)

	embed.SetFooter(fmt.Sprintf("Requested by %s", interaction.Member.User.Username), interaction.Member.User.AvatarURL(""))

	data = r.NewResponseData("").AddEmbed(embed)

	return data
}
