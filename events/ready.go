package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateWatchStatus(1, "Golang")
	fmt.Println("Ready to serve... ")
}
