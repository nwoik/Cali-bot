package main

import (
	"calibot/commands"
	"calibot/events"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("CALIBOT_TOKEN")
	session, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		log.Fatal(err)
	}

	session.AddHandler(events.Ready)
	session.AddHandler(events.MessageCreate)
	session.AddHandler(events.MemberJoin)
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

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
