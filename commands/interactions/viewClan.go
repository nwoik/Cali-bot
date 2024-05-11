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

	response := r.NewMessageResponse(ClanEmbedResponse(session, interaction, clan, members, roleInClan).InteractionResponseData)

	return response
}

func ClanEmbedResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate, clan *c.Clan, members []*m.Member, roleInClan bool) *r.Data {
	var data *r.Data

	titleEmbed := e.NewRichEmbed(fmt.Sprintf("**%s (%d/50)**", clan.Name, len(members)), "", 0xffd700)

	guildID := interaction.GuildID
	guild := GetGuild(session, guildID)

	regularMembers := FilterMembers(members, And(IsMember(session, clan), Negate(IsOfficer(session, clan)), Negate(IsLeader(clan))))
	officers := FilterMembers(members, IsOfficer(session, clan))
	leader := FilterMembers(members, IsLeader(clan))

	titleEmbed.SetThumbnail(guild.IconURL(""))
	titleEmbed.AddField("", fmt.Sprint("Clan ID: ", clan.ClanID), false)
	titleEmbed.AddField("", fmt.Sprint("**Leader: 👑 **", PrintRole(clan.LeaderRole, roleInClan)), false)
	titleEmbed.AddField("", fmt.Sprint("**Officers: 👮 **", PrintRole(clan.OfficerRole, roleInClan)), false)
	titleEmbed.AddField("", fmt.Sprint("**Members: :military_helmet: **", PrintRole(clan.MemberRole, roleInClan)), false)
	titleEmbed.AddField("**Extra Roles**", PrintExtraRoles(clan, roleInClan), false)

	leaderEmbed := e.NewRichEmbed("**Leader: **", fmt.Sprint("**👑 **", PrintRole(clan.LeaderRole, roleInClan)), 0xffd700)
	leaderEmbed.SetThumbnail(guild.IconURL(""))
	leaderEmbed = AddMemberFields(leaderEmbed, leader)

	officerEmbed := e.NewRichEmbed("**Officers: **", fmt.Sprint("**👮 **", PrintRole(clan.OfficerRole, roleInClan)), 0xffd700)
	officerEmbed.SetThumbnail(guild.IconURL(""))
	officerEmbed = AddMemberFields(officerEmbed, officers)

	memberEmbed := e.NewRichEmbed("**Members: **", fmt.Sprint("**:military_helmet: **", PrintRole(clan.OfficerRole, roleInClan)), 0xffd700)
	memberEmbed.SetThumbnail(guild.IconURL(""))
	memberEmbed = AddMemberFields(memberEmbed, regularMembers)

	// blacklistEmbed := e.NewRichEmbed("**Blacklist: **", ":no_pedestrians:", 0x000000)
	// blacklistEmbed.SetThumbnail(guild.IconURL(""))
	// blacklistEmbed = AddBlacklistFields(blacklistEmbed, clan.Blacklist)

	titleEmbed.SetFooter(fmt.Sprintf("Requested by %s", interaction.Member.User.Username), interaction.Member.User.AvatarURL(""))

	data = r.NewResponseData("")
	data.Embeds = append(data.Embeds, titleEmbed.MessageEmbed)
	data.Embeds = append(data.Embeds, leaderEmbed.MessageEmbed)
	data.Embeds = append(data.Embeds, officerEmbed.MessageEmbed)
	data.Embeds = append(data.Embeds, memberEmbed.MessageEmbed)
	// data.Embeds = append(data.Embeds, blacklistEmbed.MessageEmbed)

	return data
}
