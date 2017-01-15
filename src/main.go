package main

import (
	"app"
	"db"
	"expenses"
	"net/http"
	"routes"
)

func main() {
	env := app.New()

	db.Configure(env)
	routes.Configure(env)
	expenses.Configure(env)

	http.Handle("/", env.GetRoutes())
	http.ListenAndServe(":8080", nil)
}
