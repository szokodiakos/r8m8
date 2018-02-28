package slack

// DialogDataElement inside DialogData
type DialogDataElement struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

// DialogData inside Dialog
type DialogData struct {
	CallbackID  string              `json:"callback_id"`
	Title       string              `json:"title"`
	SubmitLabel string              `json:"submit_label"`
	Elements    []DialogDataElement `json:"elements"`
}

// Dialog in Slack's format
type Dialog struct {
	Token     string     `json:"token"`
	TriggerID string     `json:"trigger_id"`
	Dialog    DialogData `json:"dialog"`
}
