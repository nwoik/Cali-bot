package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	events "calibot/events"

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
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
