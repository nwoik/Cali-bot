package embeds

import (
	"github.com/bwmarrin/discordgo"
)

type Embed struct {
	*discordgo.MessageEmbed
}

func NewEmbed() *Embed {
	return &Embed{&discordgo.MessageEmbed{}}
}

func NewRichEmbed(title string, description string, color int) *Embed {
	richEmbed := NewEmbed().
		SetTitle(title).
		SetDescription(description).
		SetColor(color).
		SetTypeRich()

	return richEmbed
}

func (embed *Embed) AddField(name string, value string, inline bool) *Embed {
	fields := make([]*discordgo.MessageEmbedField, 0)

	fields = append(fields, createField(name, value, inline))

	embed.Fields = append(embed.Fields, fields...)

	return embed
}

func createField(name string, value string, inline bool) *discordgo.MessageEmbedField {
	return &discordgo.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	}
}

func (embed *Embed) SetImage(url string, width int, height int) *Embed {
	embed.Image = &discordgo.MessageEmbedImage{
		URL:    url,
		Width:  width,
		Height: height,
	}

	return embed
}

func (embed *Embed) SetVideo(url string, width int, height int) *Embed {
	embed.Video = &discordgo.MessageEmbedVideo{
		URL:    url,
		Width:  width,
		Height: height,
	}

	return embed
}

func (embed *Embed) SetTitle(name string) *Embed {
	embed.Title = name
	return embed
}

func (embed *Embed) SetDescription(description string) *Embed {
	embed.Description = description
	return embed
}

func (embed *Embed) SetThumbnail(url string) *Embed {

	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL: url,
	}
	return embed
}

func (embed *Embed) SetAuthor(user *discordgo.User) *Embed {
	embed.Author = &discordgo.MessageEmbedAuthor{
		Name:    user.Username,
		IconURL: user.AvatarURL(""),
	}

	return embed
}

func (embed *Embed) SetColor(color int) *Embed {
	embed.Color = color
	return embed
}

func (embed *Embed) SetFooter(text string, iconUrl string) *Embed {
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text:    text,
		IconURL: iconUrl,
	}

	return embed
}

func (embed *Embed) SetTypeRich() *Embed {
	embed.MessageEmbed.Type = discordgo.EmbedTypeRich
	return embed
}
