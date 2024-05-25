package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateStatusComplex(discordgo.UpdateStatusData{
		IdleSince: nil,
		Activities: []*discordgo.Activity{
			{
				Name:     "Pixel Gun 3D: PC Edition",
				Type:     discordgo.ActivityTypeGame,
				URL:      "https://store.steampowered.com/app/2524890/Pixel_Gun_3D_PC_Edition/",
				Instance: true,
			},
		},
		Status: "online",
		AFK:    false,
	})
	fmt.Println("Ready to serve... ")
}
