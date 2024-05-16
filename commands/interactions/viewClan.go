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

	members, _ := GetMembersWithCond(Pred("clanid", clan.ClanID))
	response := r.NewMessageResponse(ClanEmbedResponse(session, interaction, clan, members, roleInClan).InteractionResponseData)

	return response
}

func ClanEmbedResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate, clan *c.Clan, members []*m.Member, roleInClan bool) *r.Data {
	var data *r.Data

	embed := e.NewRichEmbed(fmt.Sprintf("**%s (%d/50)**", clan.Name, len(members)), "", 0xffd700)

	regularMembers, _ := GetMembersWithCond(Pred("clanid", clan.ClanID), Pred("rank", string(m.MEMBER)))
	officers, _ := GetMembersWithCond(Pred("clanid", clan.ClanID), Pred("rank", string(m.OFFICER)))
	leader, _ := GetMembersWithCond(Pred("clanid", clan.ClanID), Pred("rank", string(m.LEADER)))

	embed.SetThumbnail(GetGuild(session, interaction.GuildID).IconURL(""))
	embed.AddField("", fmt.Sprint("Clan ID: ", clan.ClanID), false)
	embed.AddField("**Extra Roles**", PrintExtraRoles(clan, roleInClan), false)
	embed.AddField("", fmt.Sprint("**Leader: 👑 **", PrintRole(clan.LeaderRole, roleInClan)), false)
	embed.AddField("", PrintMember(leader[0]), false)
	embed.AddField("", fmt.Sprint("**Officers: 👮 **", PrintRole(clan.OfficerRole, roleInClan)), false)
	embed.AddField("", PrintMembers(officers), false)

	memberEmbed := e.NewRichEmbed("", fmt.Sprint("**Members: :military_helmet: **", PrintRole(clan.MemberRole, roleInClan)), 0xd912c4)
	// memberEmbed.AddField("", fmt.Sprint("**Members: :military_helmet: **"), false)
	ms := ""
	for i, member := range regularMembers {
		ms += PrintMember(member)
		if (i%10 == 0 && i != 0) || member == regularMembers[len(regularMembers)-1] {
			memberEmbed.AddField("", ms, false)
			ms = ""
		}
	}

	blacklistEmbed := e.NewRichEmbed("", fmt.Sprint("**Blacklist :no_pedestrians: **"), 0x000)
	for i, id := range clan.Blacklist {
		ms += PingUser(id) + "\n"
		if (i%10 == 0 && i != 0) || id == clan.Blacklist[len(clan.Blacklist)-1] {
			blacklistEmbed.AddField("", ms, false)
			ms = ""
		}
	}
	// memberEmbed.AddField("Blacklist :no_pedestrians:", PrintBlacklist(clan), false)

	// embed.SetFooter(fmt.Sprintf("Requested by %s", interaction.Member.User.Username), interaction.Member.User.AvatarURL(""))

	data = r.NewResponseData("").AddEmbed(embed).AddEmbed(memberEmbed).AddEmbed(blacklistEmbed)

	return data
}

func ClanResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate, clan *c.Clan, members []*m.Member, roleInClan bool) *r.Data {
	var data *r.Data

	output := ""
	output += fmt.Sprintf("**%s (%d/50)**\n", clan.Name, len(members))

	output += fmt.Sprintf("Clan ID: %s\n", clan.ClanID)
	output += fmt.Sprintf("**Extra Roles: **%s\n", PrintExtraRoles(clan, roleInClan))

	output += fmt.Sprintf("**Members: :military_helmet: **%s\n", PrintRole(clan.MemberRole, roleInClan))
	output += fmt.Sprintf(PrintMembers(members))
	output += "\n"

	output += fmt.Sprint("**Blacklist :no_pedestrians: **\n")
	output += PrintBlacklist(clan)

	data = r.NewResponseData(output)

	return data
}
