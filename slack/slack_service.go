package slack

import (
	"net/url"

	"github.com/szokodiakos/r8m8/slack/errors"
)

// Service interface
type Service interface {
	ParseRequestValues(values string) (RequestValues, error)
	CreateMessageResponse(text string) MessageResponse
	VerifyToken(values string) error
}

type slackService struct {
	verificationToken string
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
	userID := parsedValues.Get("user_id")
	userName := parsedValues.Get("user_name")

	requestValues = RequestValues{
		Text:        text,
		Token:       token,
		TeamID:      teamID,
		TeamDomain:  teamDomain,
		ChannelID:   channelID,
		ChannelName: channelName,
		UserID:      userID,
		UserName:    userName,
	}
	return requestValues, nil
}

func (s *slackService) CreateMessageResponse(text string) MessageResponse {
	return MessageResponse{
		Text: text,
	}
}

func (s *slackService) VerifyToken(values string) error {
	requestValues, err := s.ParseRequestValues(values)
	if err != nil {
		return err
	}

	token := requestValues.Token
	if token != s.verificationToken {
		return &errors.InvalidVerificationTokenError{}
	}

	return nil
}

// NewService factory
func NewService(verificationToken string) Service {
	return &slackService{
		verificationToken: verificationToken,
	}
}
