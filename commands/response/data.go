package response

import (
	embeds "calibot/components/embeds"

	"github.com/bwmarrin/discordgo"
)

type Data struct {
	*discordgo.InteractionResponseData
}

func NewData() *Data {
	return &Data{&discordgo.InteractionResponseData{}}
}

func NewResponseData(content string) *Data {
	responseData := NewData().
		SetContent(content)

	return responseData
}

func (data *Data) SetContent(content string) *Data {
	data.Content = content

	return data
}

func (data *Data) AddEmbed(embed *embeds.Embed) *Data {
	embeds := make([]*discordgo.MessageEmbed, 0)

	embeds = append(embeds, embed.MessageEmbed)

	data.Embeds = append(data.Embeds, embeds...)

	return data
}

// func (data *Data) AddComponent(content string) *Data {
// 	data.Content = content

// 	return data
// }
