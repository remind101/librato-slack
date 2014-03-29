package services

import(
	"fmt"
	"net/http"
	"github.com/remind101/slack-hooks/slack"
)

func Papertrail(res http.ResponseWriter, req *http.Request) string {
	account := req.FormValue("account")
	token   := req.FormValue("token")
	//channel := req.FormValue("channel")

	payload := slack.Payload{
		"Papertrail",
		"Search",
		nil,
	}

	response, err := slack.Post(account, token, payload)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)

	return "Papertrail"
}
