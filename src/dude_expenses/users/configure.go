package users

import (
	"dude_expenses/app"
)

func Configure(env *app.Env) {
	router := env.GetRouter()
	router.Handle("/users", app.Handle(Create(env))).Methods("POST")
	router.Handle("/users/authenticate", app.Handle(Authenticate(env))).Methods("POST")
}
