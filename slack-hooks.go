package main

import(
	"net/http"
	"github.com/codegangsta/martini"
	"github.com/remind101/slack-hooks/services"
	"github.com/remind101/slack-hooks/slack"
)

func main() {
	m := martini.Classic()

	m.Use(postMessage)

	m.Post("/librato",    services.Librato)
	m.Post("/papertrail", services.Papertrail)

	m.Run()
}

func postMessage(res http.ResponseWriter, req *http.Request, c martini.Context) {
	message := slack.Message{}
	c.Map(&message)
	c.Next()

	account := req.FormValue("account")
	token   := req.FormValue("token")
	slack.Post(account, token, message)
}
