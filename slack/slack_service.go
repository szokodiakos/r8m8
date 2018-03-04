package slack

import "net/url"

// Service interface
type Service interface {
	ParseRequestValues(values string) (RequestValues, error)
	CreateMessageResponse(text string) MessageResponse
}

type slackService struct {
}

func (ss *slackService) ParseRequestValues(values string) (RequestValues, error) {
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

	requestValues = RequestValues{
		Text:   text,
		Token:  token,
		TeamID: teamID,
	}
	return requestValues, nil
}

func (ss *slackService) CreateMessageResponse(text string) MessageResponse {
	return MessageResponse{
		Text: text,
	}
}

// NewService factory
func NewService() Service {
	return &slackService{}
}
