package services

import(
	"github.com/remind101/slack-hooks/slack"
)

func Papertrail(message *slack.Message) {
	message.Username = "Papertrail"
	message.Text	 = "Search"
}
