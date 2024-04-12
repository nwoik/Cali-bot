package main

import (
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
	session.AddHandler(interactionCreate)
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}

	err = registerCommand(session)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Handle interaction type
	if i.Type == discordgo.InteractionApplicationCommand {
		// Handle the "/hello" command
		if i.ApplicationCommandData().Name == "hello" {
			// Respond to the command with "hello"
			response := &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "hello",
				},
			}
			_ = s.InteractionRespond(i.Interaction, response)
		}
	}
}

func registerCommand(s *discordgo.Session) error {
	// Define the command structure
	// command := commands.NewChatApplicationCommandCommand("hello", "say hello")
	command := &discordgo.ApplicationCommand{
		Name:        "hello",
		Description: "Say hello",
		Type:        discordgo.ChatApplicationCommand,
	}

	// Register the command globally
	_, err := s.ApplicationCommandCreate(s.State.User.ID, s.State.Application.GuildID, command)
	if err != nil {
		return err
	}

	return nil
}
