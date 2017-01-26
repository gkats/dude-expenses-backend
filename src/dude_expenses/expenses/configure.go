package expenses

import (
	"dude_expenses/app"
)

func Configure(env *app.Env) {
	router := env.GetRouter()
	router.Handle("/expenses", app.Handle(Index(env))).Methods("GET")
	router.Handle("/expenses", app.Handle(Create(env))).Methods("POST")
}
