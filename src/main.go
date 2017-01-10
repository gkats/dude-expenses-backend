package main

import (
	"app"
	"db"
)

func main() {
	env := app.New()

	db.Configure(env)
}