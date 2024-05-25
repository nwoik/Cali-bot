package events

import (
	"calibot/commands/interactions"
	"calibot/globals"
	"fmt"
	"log"

	c "github.com/nwoik/calibotapi/model/clan"

	"github.com/bwmarrin/discordgo"
)

func MemberBan(session *discordgo.Session, member *discordgo.GuildBanAdd) {
	fmt.Printf("Member banned from the server: %s\n", member.User.Username)

	client := globals.CLIENT

	clanCollection := client.Database("calibot").Collection("clan")
	clanRepo := c.NewClanRepo(clanCollection)

	guildid := member.GuildID

	clan, err := interactions.GetClan(guildid)

	if err != nil {
		log.Println("Clan not found")
	}

	clan = Blacklist(clan, member)
	clanRepo.Update(clan)

}

func Blacklist(clan *c.Clan, member *discordgo.GuildBanAdd) *c.Clan {
	userid := member.User.ID

	if !interactions.IsBlacklisted(clan, userid) {
		clan.Blacklist = append(clan.Blacklist, userid)
	}

	return clan
}
