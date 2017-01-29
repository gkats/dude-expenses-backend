package main

import (
	"dude_expenses/app"
	"dude_expenses/db"
	"dude_expenses/expenses"
	"dude_expenses/users"
	"flag"
	"net/http"
	"strconv"
)

func main() {
	var (
		authSecret = flag.String(
			"auth_secret",
			"c7b5c187c13400dac",
			"Secret key for encoding/decoding JWTs.",
		)
		port  = flag.Int("port", 8080, "The HTTP port the server will listen to.")
		dbUrl = flag.String(
			"db_url",
			"dude_expenses:dudeExpen$es123@localhost:5432/dude_expenses_development",
			"The Postgresql database URL. Should be in 'user:password@host:port/database' format.",
		)
	)
	flag.Parse()

	env := app.New(*authSecret, *dbUrl)

	db.Configure(env)
	users.Configure(env)
	expenses.Configure(env)

	http.Handle("/", env.GetRouter())
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}
