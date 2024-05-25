package interactions

import (
	r "calibot/commands/response"
	"calibot/components/button"
	e "calibot/components/embeds"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/model/clan"
)

func ViewBlacklist(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	var roleInClan bool
	var clanid string

	message := interaction.Message
	embed := message.Embeds[0]

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

	response := r.NewMessageResponse(BlackListEmbedResponse(session, interaction, clan, roleInClan).InteractionResponseData)

	return response
}

func BlackListEmbedResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate, clan *c.Clan, roleInClan bool) *r.Data {
	embed := e.NewRichEmbed(fmt.Sprintf("%s: %d", clan.Name, len(clan.Blacklist)), "", 0x000)
	embed.AddField("", fmt.Sprint("Clan ID: ", clan.ClanID), false)
	embed.AddField("", "**Blacklist :no_pedestrians: **", false)

	ms := ""

	for i, id := range clan.Blacklist {
		ms += PingUser(id) + "\n"
		if (i%5 == 0 && i != 0) || id == clan.Blacklist[len(clan.Blacklist)-1] {
			embed.AddField("", ms, false)
			ms = ""
		}
	}

	data := r.NewResponseData("").AddEmbed(embed)

	actionRow := e.NewActionRow()
	actionRow2 := e.NewActionRow()

	previousButton := button.NewBasicButton("Previous", "clan_previous_button", discordgo.PrimaryButton, true)
	homeButton := button.NewBasicButton("Home", "clan_home_button", discordgo.SecondaryButton, false)
	nextButton := button.NewBasicButton("Next", "clan_next_button", discordgo.PrimaryButton, true)
	blacklistButton := button.NewBasicButton("Blacklist", "blacklist_button", discordgo.DangerButton, true)

	actionRow.Components = append(actionRow.Components, previousButton)
	actionRow.Components = append(actionRow.Components, homeButton)
	actionRow.Components = append(actionRow.Components, nextButton)
	actionRow2.Components = append(actionRow2.Components, blacklistButton)

	data.Components = append(data.Components, actionRow)
	data.Components = append(data.Components, actionRow2)

	return data
}
