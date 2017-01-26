package users

import (
	"dude_expenses/app"
	"dude_expenses/log"
)

func Configure(env *app.Env) {
	router := env.GetRouter()
	router.Handle("/users", log.WithLogging(env, app.Handle(Create(env)))).Methods("POST")
	router.Handle("/users/authenticate", log.WithLogging(env, app.Handle(Authenticate(env)))).Methods("POST")
}
