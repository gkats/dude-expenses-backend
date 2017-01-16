package expenses

import (
	"app"
	"app/handler"
)

func Configure(env *app.Env) {
	router := env.GetRouter()
	router.Handle("/expenses", handler.AppHandler(env, Index)).Methods("GET")
	router.Handle("/expenses", handler.AppHandler(env, Create)).Methods("POST")
}
