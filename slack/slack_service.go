package slack

import "net/url"

// Service interface
type Service interface {
	ParseRequestValues(values string) RequestValues
}

type slackService struct {
}

func (ss *slackService) ParseRequestValues(values string) RequestValues {
	parsedValues, _ := url.ParseQuery(values)

	rawText := parsedValues.Get("text")
	text, _ := url.QueryUnescape(rawText)
	token := parsedValues.Get("token")
	teamID := parsedValues.Get("team_id")

	return RequestValues{
		Text:   text,
		Token:  token,
		TeamID: teamID,
	}
}

// NewService factory
func NewService() Service {
	return &slackService{}
}
