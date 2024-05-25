package main

import (
	"calibot/commands"
	events "calibot/events"
	"calibot/globals"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	globals.InitConfig()

	session, err := discordgo.New(fmt.Sprintf("Bot %s", globals.TOKEN))
	if err != nil {
		log.Fatal(err)
	}

	session.AddHandler(events.Ready)
	session.AddHandler(events.MessageCreate)
	session.AddHandler(events.MemberJoin)
	session.AddHandler(events.MemberBan)
	session.AddHandler(events.MemberLeave)
	session.AddHandler(commands.InteractionCreate)
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}

	err = commands.RegisterCommand(session)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()
	defer globals.CLIENT.Disconnect(context.Background())

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
