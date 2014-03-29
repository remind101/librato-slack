package main

import(
	"net/http"
	"github.com/codegangsta/martini"
	"github.com/remind101/slack-hooks/services"
	"github.com/remind101/slack-hooks/slack"
)

func main() {
	m := martini.Classic()

	m.Use(func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		payload := slack.Payload{}
		c.Map(&payload)
		c.Next()

		account := req.FormValue("account")
		token   := req.FormValue("token")
		slack.Post(account, token, payload)
	})

	m.Post("/librato",    services.Librato)
	m.Post("/papertrail", services.Papertrail)

	m.Run()
}
