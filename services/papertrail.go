package services

import(
	"github.com/remind101/slack-hooks/slack"
)

func Papertrail(payload *slack.Payload) {
	payload.Username = "Papertrail"
	payload.Text	 = "Search"
}
