package main

import(
	"net/http"
	"github.com/codegangsta/martini"
	"github.com/remind101/slack-hooks/slack"
	"github.com/remind101/slack-hooks/services"
	"github.com/remind101/slack-hooks/services/papertrail"
)

func main() {
	m := martini.Classic()

	handlers := map[string]services.Service{
		"/papertrail": &papertrail.Service{},
	}

	for path, handler := range handlers {
		m.Post(path, func(res http.ResponseWriter, req *http.Request, c martini.Context) {
			message := handler.Handle(handler.Parse())

			account := req.FormValue("account")
			token   := req.FormValue("token")

			go slack.Post(account, token, message)
		})
	}

	m.Run()
}
