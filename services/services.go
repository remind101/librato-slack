package services

import(
	"github.com/remind101/slack-hooks/slack"
)

type Service interface {
	Parse(string) interface{}
	Handle(interface{}) *slack.Message
}
