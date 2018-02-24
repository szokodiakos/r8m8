package leaderboard

import (
	"bytes"
	"fmt"
	"net/http"
)

type slackPrinter struct {
}

// Print content via Slack webhook
func (p *slackPrinter) Print() {
	reader := bytes.NewReader([]byte(`{ "text": "Hello from Go" }`))
	_, err := http.Post("https://hooks.slack.com/services/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", "application/json", reader)
	if err != nil {
		fmt.Println("Err", err)
	}
}

// NewSlackPrinter creates a slackPrinter
func NewSlackPrinter() Printer {
	return &slackPrinter{}
}
