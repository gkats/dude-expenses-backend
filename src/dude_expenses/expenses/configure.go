package expenses

import (
	"dude_expenses/app"
	"dude_expenses/auth"
)

func Configure(env *app.Env) {
	router := env.GetRouter()
	router.Handle("/expenses", app.Handle(auth.WithAuth(env, Index(env)))).Methods("GET")
	router.Handle("/expenses", app.Handle(auth.WithAuth(env, Create(env)))).Methods("POST")
}
