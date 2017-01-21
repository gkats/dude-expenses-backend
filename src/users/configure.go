package users

import (
	"app"
	"app/handler"
)

func Configure(env *app.Env) {
	router := env.GetRouter()
	router.Handle("/users", handler.AppHandler(env, Create)).Methods("POST")
	router.Handle("/authenticate", handler.AppHandler(env, Authenticate)).Methods("POST")
}
