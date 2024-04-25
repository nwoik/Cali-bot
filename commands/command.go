package commands

import (
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	*discordgo.ApplicationCommand
}

func NewCommand() *Command {
	return &Command{&discordgo.ApplicationCommand{}}
}

func newApplicationCommand(name string, description string, commandType discordgo.ApplicationCommandType) *Command {
	botCommand := NewCommand().
		SetName(name).
		SetDescription(description).
		SetApplicationCommandType(commandType)

	return botCommand
}

func NewChatApplicationCommand(name string, description string) *Command {
	return newApplicationCommand(name, description,
		discordgo.ChatApplicationCommand)
}

func NewUserApplicationCommand(name string, description string) *Command {
	return newApplicationCommand(name, description,
		discordgo.UserApplicationCommand)

}

func NewMessageApplicationCommand(name string, description string) *Command {
	return newApplicationCommand(name, description,
		discordgo.MessageApplicationCommand)
}

func (command *Command) SetName(name string) *Command {
	command.Name = name

	return command
}

func (command *Command) SetDescription(description string) *Command {
	command.Description = description

	return command
}

func (command *Command) SetDefaultMemberPermissions(permission int64) *Command {
	command.DefaultMemberPermissions = &permission

	return command
}

func (command *Command) SetApplicationCommandType(commandType discordgo.ApplicationCommandType) *Command {
	command.Type = discordgo.ApplicationCommandType(commandType)

	return command
}

func (command *Command) SetOptions(options []*discordgo.ApplicationCommandOption) *Command {
	command.Options = options

	return command
}
