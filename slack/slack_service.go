package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Service Slack
type Service interface {
	OpenMatchDialog(values string)
}

type service struct {
}

func (s *service) OpenMatchDialog(values string) {
	fmt.Println(values)
	parsedValues, _ := url.ParseQuery(values)
	triggerID := parsedValues.Get("trigger_id")
	dialog := s.getMatchDialog(triggerID, triggerID)
	token := "xoxp-26143804727-26136563206-321504853783-2a730ec1c54f5f7ef49a3a98d13f031c"
	s.openDialog(token, dialog)
}

func (s *service) openDialog(token string, dialog Dialog) error {
	marshalledDialog, _ := json.Marshal(dialog)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://slack.com/api/dialog.open", bytes.NewReader(marshalledDialog))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	_, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// GetMatchDialog returns a Slack dialog of Match data
func (s *service) getMatchDialog(triggerID string, callbackID string) Dialog {
	return Dialog{
		TriggerID: triggerID,
		Dialog: DialogData{
			CallbackID:  callbackID,
			Title:       "Match Details",
			SubmitLabel: "Add",
			Elements: []DialogDataElement{
				DialogDataElement{
					Type:  "text",
					Label: "Winner Team Members",
					Name:  "winner_team_members",
				},
				DialogDataElement{
					Type:  "text",
					Label: "Loser Team Members",
					Name:  "loser_team_members",
				},
			},
		},
	}
}

// NewService creates a service
func NewService() Service {
	return &service{}
}
