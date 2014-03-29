package slack

import(
	"encoding/json"
	"net/http"
	"bytes"
	//"strings"
)

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type Attachment struct {
	Fields []Field `json:"fields"`
}

type Payload struct {
	Username    string	 `json:"username"`
	Text	    string	 `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

func Post(account string, token string, payload Payload) (*http.Response, error) {
	url := "https://" + account + ".slack.com/services/hooks/incoming-webhook?token=" + token

	json, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	return http.Post(url, "application/json", bytes.NewBuffer(json))
}
