package emoji

import (
	"github.com/bwmarrin/discordgo"
)

type Emoji struct {
	*discordgo.ComponentEmoji
}

func NewEmoji() *Emoji {
	return &Emoji{&discordgo.ComponentEmoji{}}
}

func NewBasicEmoji(name string) *Emoji {
	emoji := NewEmoji().
		SetName(name).
		SetID("").
		SetAnimated(false)

	return emoji
}

func NewComponentEmoji(name string, id string, animated bool) *Emoji {
	emoji := NewEmoji().
		SetName(name).
		SetID(id).
		SetAnimated(animated)

	return emoji
}

func (emoji *Emoji) SetName(name string) *Emoji {
	emoji.Name = name
	return emoji
}

func (emoji *Emoji) SetID(id string) *Emoji {
	emoji.ID = id
	return emoji
}

func (emoji *Emoji) SetAnimated(animated bool) *Emoji {
	emoji.Animated = animated

	return emoji
}
