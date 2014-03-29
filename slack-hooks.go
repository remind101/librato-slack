package main

import(
	"github.com/codegangsta/martini"
	"github.com/remind101/slack-hooks/services"
)

func main() {
	m := martini.Classic()

	m.Post("/librato",    services.Librato)
	m.Post("/papertrail", services.Papertrail)

	m.Run()
}
