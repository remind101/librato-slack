package papertrail

import (
	"encoding/json"
	"github.com/remind101/slack-hooks/slack"
)

type Event struct {
	Hostname string `json:"hostname"`
	Message  string `json:"message"`
}

type Payload struct {
	Events []Event `json:"events"`
}

type Service struct {
}

func (s *Service) Parse(input string) interface{} {
	payload := &Payload{}

	if err := json.Unmarshal([]byte(input), payload); err != nil {
		panic(err)
	}

	return payload
}

func (s *Service) Handle(payload interface{}) *slack.Message {
	return &slack.Message{
		Username: "Papertrail",
		Text:     "Search",
	}
}
