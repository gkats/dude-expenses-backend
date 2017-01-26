package expenses

import (
	"dude_expenses/app"
	"dude_expenses/auth"
	"dude_expenses/log"
)

func Configure(env *app.Env) {
	router := env.GetRouter()
	router.Handle("/expenses", log.WithLogging(env, app.Handle(auth.WithAuth(env, Index(env))))).Methods("GET")
	router.Handle("/expenses", log.WithLogging(env, app.Handle(auth.WithAuth(env, Create(env))))).Methods("POST")
}
