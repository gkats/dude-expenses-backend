package main

import (
	"dude_expenses/app"
	"dude_expenses/db"
	"dude_expenses/expenses"
	"dude_expenses/users"
	"net/http"
)

func main() {
	env := app.New()

	db.Configure(env)
	users.Configure(env)
	expenses.Configure(env)

	http.Handle("/", env.GetRouter())
	http.ListenAndServe(":8080", nil)
}
