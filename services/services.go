package services

import(
	"github.com/remind101/slack-hooks/slack"
)

type Service interface {
	Parse() interface{}
	Handle(*interface{}) *slack.Message
}
