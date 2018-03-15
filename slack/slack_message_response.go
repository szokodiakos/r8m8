package slack

// MessageResponse struct
type MessageResponse struct {
	Text         string `json:"text"`
	ResponseType string `json:"response_type"`
}

// CreateDirectResponse func
func CreateDirectResponse(text string) MessageResponse {
	return MessageResponse{
		Text:         text,
		ResponseType: "ephemeral",
	}
}

// CreateChannelResponse func
func CreateChannelResponse(text string) MessageResponse {
	return MessageResponse{
		Text:         text,
		ResponseType: "in_channel",
	}
}
