package interactions

import (
	r "calibot/commands/responses"
	e "calibot/embeds"
	"fmt"

	"github.com/bwmarrin/discordgo"
	c "github.com/nwoik/calibotapi/clan"
	m "github.com/nwoik/calibotapi/member"
)

func ViewMember(session *discordgo.Session, interaction *discordgo.InteractionCreate) *r.Response {
	members := m.Open("./resources/members.json")

	args := interaction.ApplicationCommandData().Options
	var member *m.Member

	if len(args) != 0 {
		user := GetArgument(args, "member").UserValue(session)
		member = GetMember(members, user.ID)
	} else {
		member = GetMember(members, interaction.Member.User.ID)
	}

	response := r.NewMessageResponse(EmbedResponse(session, interaction, member).InteractionResponseData)

	return response
}

func EmbedResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate, member *m.Member) *r.Data {
	var data *r.Data

	if member == nil {
		data = r.NewResponseData("User is not registered with the bot.")
	} else {
		guildID := interaction.GuildID
		memberID := interaction.Member.User.ID

		// Get the member's information
		user, err := session.GuildMember(guildID, memberID)
		if err != nil {
			fmt.Println("Error retrieving member information:", err)
		}

		embed := e.NewRichEmbed(member.Nick, "User Info", 0x08d052c)
		embed.SetThumbnail(user.AvatarURL(""))

		embed.AddField("**IGN: **", member.IGN, false)
		embed.AddField("**ID: **", member.IGID, false)

		if member.ClanID != "" {
			clans := c.Open("./resources/clan.json")
			clan := GetClanByClanID(clans, member.ClanID)
			embed.AddField("**Clan: **", clan.Name, true)
		}

		embed.SetFooter(fmt.Sprintf("Requested by %s", interaction.Member.User.Username), interaction.Member.User.AvatarURL(""))
		data = r.NewResponseData("").AddEmbed(embed)
	}

	return data
}
