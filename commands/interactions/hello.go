package interactions

import (
	responses "calibot/commands/responses"
)

func Hello() *responses.Response {
	data := responseData()
	response := responses.NewMessageResponse(data.InteractionResponseData)

	return response
}

func responseData() *responses.Data {
	data := responses.NewResponseData("Hello there!")

	return data
}
