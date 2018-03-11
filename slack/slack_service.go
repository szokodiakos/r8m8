package slack

import "net/url"

// Service interface
type Service interface {
	ParseRequestValues(values string) (RequestValues, error)
	CreateMessageResponse(text string) MessageResponse
}

type slackService struct {
}

func (s *slackService) ParseRequestValues(values string) (RequestValues, error) {
	var requestValues RequestValues
	parsedValues, err := url.ParseQuery(values)
	if err != nil {
		return requestValues, err
	}

	rawText := parsedValues.Get("text")
	text, err := url.QueryUnescape(rawText)
	if err != nil {
		return requestValues, err
	}

	token := parsedValues.Get("token")
	teamID := parsedValues.Get("team_id")
	teamDomain := parsedValues.Get("team_domain")
	channelID := parsedValues.Get("channel_id")
	channelName := parsedValues.Get("channel_name")

	requestValues = RequestValues{
		Text:        text,
		Token:       token,
		TeamID:      teamID,
		TeamDomain:  teamDomain,
		ChannelID:   channelID,
		ChannelName: channelName,
	}
	return requestValues, nil
}

func (s *slackService) CreateMessageResponse(text string) MessageResponse {
	return MessageResponse{
		Text: text,
	}
}

// NewService factory
func NewService() Service {
	return &slackService{}
}
