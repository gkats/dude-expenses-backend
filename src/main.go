package main

import (
	"net/http"
	"app"
	"db"
	"expenses"
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
