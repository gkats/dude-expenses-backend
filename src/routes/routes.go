package routes

import (
	"app"
	"github.com/gorilla/mux"
)

func Configure(env *app.Env) {
	env.SetRoutes(mux.NewRouter())
}
