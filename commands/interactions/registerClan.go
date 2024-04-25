package interactions

import (
	responses "calibot/commands/responses"

	c "github.com/nwoik/calibotapi/clan"

	"github.com/bwmarrin/discordgo"
)

func RegisterClan(session *discordgo.Session, interaction *discordgo.InteractionCreate) *responses.Response {
	clans := c.Open("./resources/clan.json")

}
