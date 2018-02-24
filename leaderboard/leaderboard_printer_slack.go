package leaderboard

import (
	"bytes"
	"fmt"
	"net/http"
)

type slackPrinter struct {
	webhook string
}

// Print content via Slack webhook
func (p *slackPrinter) Print() {
	reader := bytes.NewReader([]byte(`{ "text": "Hello from Go" }`))
	_, err := http.Post(p.webhook, "application/json", reader)
	if err != nil {
		fmt.Println("Err", err)
	}
}

// NewSlackPrinter creates a slackPrinter
func NewSlackPrinter(webhook string) Printer {
	return &slackPrinter{
		webhook,
	}
}
