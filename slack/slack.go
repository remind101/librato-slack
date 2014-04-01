package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
	//"strings"
)

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type Attachment struct {
	Fields []Field `json:"fields"`
}

type Message struct {
	Username    string       `json:"username"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

func Post(account string, token string, message *Message) (*http.Response, error) {
	url := "https://" + account + ".slack.com/services/hooks/incoming-webhook?token=" + token

	json, err := json.Marshal(message)

	if err != nil {
		return nil, err
	}

	return http.Post(url, "application/json", bytes.NewBuffer(json))
}
