package expenses

import (
	"app"
)

func Configure(env *app.Env) {
	routes := env.GetRoutes()
	routes.Handle("/expenses", Index(env)).Methods("GET")
	routes.Handle("/expenses", Create(env)).Methods("POST")
}
