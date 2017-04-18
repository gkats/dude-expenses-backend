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
	router.Handle("/expenses/tags", log.WithLogging(env, app.Handle(auth.WithAuth(env, GetTags(env))))).Methods("GET")
	router.Handle("/expenses/{id:[0-9]+}", log.WithLogging(env, app.Handle(auth.WithAuth(env, Show(env))))).Methods("GET")
	router.Handle("/expenses/{id:[0-9]+}", log.WithLogging(env, app.Handle(auth.WithAuth(env, Update(env))))).Methods("PATCH")
}
