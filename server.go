package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/api", func() string {
		return "Return api"
	})	
	m.Get("/application", func() string {
		return "List applications"
	})
	m.Post("/application/new", func() string {
		return "New application"
	})
	m.Get("/application/:name", func(params martini.Params) string {
		return "Detail application " + params["name"]
	})
	m.Put("/application/:name/edit", func(params martini.Params) string {
		return "Edit application " + params["name"]
	})
	m.Delete("/application/:name/delete", func(params martini.Params) string {
		return "Edit application " + params["name"]
	})
	m.Run()
}
