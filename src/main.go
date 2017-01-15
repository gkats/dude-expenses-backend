package main

import (
	"app"
	"db"
	"expenses"
	"net/http"
)

func main() {
	env := app.New()

	db.Configure(env)
	expenses.Configure(env)

	http.Handle("/", env.GetRouter())
	http.ListenAndServe(":8080", nil)
}
