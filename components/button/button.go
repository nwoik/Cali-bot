package button

import (
	"github.com/bwmarrin/discordgo"
)

type Button struct {
	*discordgo.Button
}

func NewButton() *Button {
	return &Button{&discordgo.Button{}}
}

func NewEmojiButton(label string, style discordgo.ButtonStyle, disabled bool, emoji *discordgo.ComponentEmoji) *Button {

	button := NewButton().
		SetLabel(label).
		SetStyle(style).
		SetDisabled(disabled).
		SetEmoji(emoji)

	return button
}

func (button *Button) SetLabel(label string) *Button {
	button.Label = label

	return button
}

func (button *Button) SetStyle(style discordgo.ButtonStyle) *Button {
	button.Style = style

	return button
}

func (button *Button) SetDisabled(disabled bool) *Button {
	button.Disabled = disabled

	return button
}

func (button *Button) SetEmoji(emoji *discordgo.ComponentEmoji) *Button {
	button.Emoji = emoji

	return button
}
