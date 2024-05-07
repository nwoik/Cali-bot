package response

import "github.com/bwmarrin/discordgo"

type Response struct {
	*discordgo.InteractionResponse
}

func NewResponse() *Response {
	return &Response{&discordgo.InteractionResponse{}}
}

func NewInteractionResponse(data *discordgo.InteractionResponseData, responseType discordgo.InteractionResponseType) *Response {
	response := NewResponse().
		SetInteractionResponseData(data).
		SetApplicationCommandType(responseType)

	return response
}

func NewMessageResponse(data *discordgo.InteractionResponseData) *Response {
	response := NewResponse().
		SetInteractionResponseData(data).
		SetApplicationCommandType(discordgo.InteractionResponseChannelMessageWithSource)

	return response
}

func (response *Response) SetInteractionResponseData(data *discordgo.InteractionResponseData) *Response {
	response.Data = data

	return response
}

func (response *Response) SetApplicationCommandType(responseType discordgo.InteractionResponseType) *Response {
	response.Type = discordgo.InteractionResponseType(responseType)

	return response
}
