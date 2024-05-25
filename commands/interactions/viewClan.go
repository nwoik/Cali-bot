package interactions

import (
	r "calibot/commands/response"
	"calibot/components/button"
	e "calibot/components/embeds"
	"fmt"
	"strconv"
	"strings"

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
	response := r.NewMessageResponse(ClanEmbedResponse(session, interaction, clan, members, roleInClan, 0).InteractionResponseData)

	return response
}

func HomePage(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	var roleInClan bool
	var clanid string

	message := interaction.Message
	embed := message.Embeds[0]
	pageNum := 0

	for _, field := range embed.Fields {
		if strings.Contains(field.Value, "Clan ID") {
			clanid = strings.ReplaceAll(strings.Split(field.Value, ":")[1], " ", "")
		}
	}

	if clanid == interaction.GuildID {
		roleInClan = true
	} else {
		roleInClan = false
	}

	clan, err := GetClan(clanid)

	if err != nil {
		return r.NewMessageResponse(r.ClanNotRegisteredWithGuild().InteractionResponseData)
	}

	members, _ := GetMembersWithCond(Pred("clanid", clan.ClanID))
	response := r.NewMessageResponse(ClanEmbedResponse(session, interaction, clan, members, roleInClan, pageNum).InteractionResponseData)

	return response
}

func IncPage(session *discordgo.Session, interaction *discordgo.InteractionCreate, inc int) *r.Response {
	var roleInClan bool
	var clanid string

	message := interaction.Message
	embed := message.Embeds[0]
	page := strings.ReplaceAll(strings.Split(embed.Footer.Text, ":")[1], " ", "")
	pageNum, err := strconv.ParseInt(page, 10, 64)
	pageNum += int64(inc)

	if pageNum < 0 {
		pageNum = 0
	} else if pageNum > 8 {
		pageNum = 8
	}

	if err != nil {
		return r.NewMessageResponse(r.NewResponseData("Error changing page").InteractionResponseData)
	}

	for _, field := range embed.Fields {
		if strings.Contains(field.Value, "Clan ID") {
			clanid = strings.ReplaceAll(strings.Split(field.Value, ":")[1], " ", "")
		}
	}

	if clanid == interaction.GuildID {
		roleInClan = true
	} else {
		roleInClan = false
	}

	clan, err := GetClan(clanid)

	if err != nil {
		return r.NewMessageResponse(r.ClanNotRegisteredWithGuild().InteractionResponseData)
	}

	members, _ := GetMembersWithCond(Pred("clanid", clan.ClanID))
	response := r.NewMessageResponse(ClanEmbedResponse(session, interaction, clan, members, roleInClan, int(pageNum)).InteractionResponseData)

	return response
}

func ClanEmbedResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate, clan *c.Clan, members []*m.Member, roleInClan bool, page int) *r.Data {
	var data *r.Data
	var beginning, end bool

	embed := e.NewRichEmbed(fmt.Sprintf("**%s (%d/50)**", clan.Name, len(members)), "", 0xffd700)

	regularMembers, _ := GetMembersWithCond(Pred("clanid", clan.ClanID), Pred("rank", string(m.MEMBER)))
	officers, _ := GetMembersWithCond(Pred("clanid", clan.ClanID), Pred("rank", string(m.OFFICER)))
	leader, _ := GetMembersWithCond(Pred("clanid", clan.ClanID), Pred("rank", string(m.LEADER)))

	embed.SetThumbnail(GetGuild(session, interaction.GuildID).IconURL(""))
	embed.AddField("", fmt.Sprint("Clan ID: ", clan.ClanID), false)

	if page == 0 {
		embed.AddField("**Extra Roles**", PrintExtraRoles(clan, roleInClan), false)
		embed.AddField("", fmt.Sprint("**Leader: 👑 **", PrintRole(clan.LeaderRole, roleInClan)), false)
		embed.AddField("", PrintMembers(leader), false)
		embed.AddField("", fmt.Sprint("**Officers: 👮 **", PrintRole(clan.OfficerRole, roleInClan)), false)
		embed.AddField("", PrintMembers(officers), false)
		beginning = true
	} else {
		beginning, end = false, false

		start := (page - 1) * 5
		ms := ""
		for i := 0; i < 5; i++ {
			if (start + i) < len(regularMembers) {
				member := regularMembers[start+i]
				ms += PrintMember(member)
			}
		}

		if ms != "" {
			embed.SetColor(0xd912c4)
			embed.AddField("", fmt.Sprint("**Members: :military_helmet: **", PrintRole(clan.MemberRole, roleInClan)), false)
			embed.AddField("", ms, false)
		} else {
			end = true
			embed.AddField("End", ":end:", false)
		}
	}

	embed.SetFooter(fmt.Sprintf("Page : %d", page), "")

	data = r.NewResponseData("").AddEmbed(embed)
	data.CustomID = interaction.ID

	actionRow := e.NewActionRow()
	actionRow2 := e.NewActionRow()

	previousButton := button.NewBasicButton("Previous", "clan_previous_button", discordgo.PrimaryButton, beginning)
	homeButton := button.NewBasicButton("Home", "clan_home_button", discordgo.SecondaryButton, false)
	nextButton := button.NewBasicButton("Next", "clan_next_button", discordgo.PrimaryButton, end)
	blacklistButton := button.NewBasicButton("Blacklist", "blacklist_button", discordgo.DangerButton, end)

	actionRow.Components = append(actionRow.Components, previousButton)
	actionRow.Components = append(actionRow.Components, homeButton)
	actionRow.Components = append(actionRow.Components, nextButton)
	actionRow2.Components = append(actionRow2.Components, blacklistButton)

	data.Components = append(data.Components, actionRow)
	data.Components = append(data.Components, actionRow2)

	return data
}
