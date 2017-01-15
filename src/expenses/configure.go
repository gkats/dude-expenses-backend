package expenses

import (
	"app"
)

func Configure(env *app.Env) {
	router := env.GetRouter()
	router.Handle("/expenses", Index(env)).Methods("GET")
	router.Handle("/expenses", Create(env)).Methods("POST")
}
