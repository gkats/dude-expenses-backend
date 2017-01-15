package main

import (
	"app"
	"db"
	"expenses"
	"net/http"
	"users"
)

func main() {
	env := app.New()

	db.Configure(env)
	users.Configure(env)
	expenses.Configure(env)

	http.Handle("/", env.GetRouter())
	http.ListenAndServe(":8080", nil)
}
