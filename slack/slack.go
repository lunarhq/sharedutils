package slack

import (
	"fmt"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/lunarhq/sharedutils/env"
)

var (
	Endpoint = env.Get("SLACK_WEBHOOK_URL", "http://localhost:8080")
)

type Field struct {
	Title string
	Value string
}

func Post(title string, items ...Field) {
	att := slack.Attachment{}
	for _, item := range items {
		att.AddField(slack.Field{Title: item.Title, Value: item.Value})
	}

	payload := slack.Payload{Text: title, Attachments: []slack.Attachment{att}}
	if err := slack.Send(Endpoint, "", payload); len(err) > 0 {
		fmt.Printf("Err sending to slack: %s\n", err)
	}
}
