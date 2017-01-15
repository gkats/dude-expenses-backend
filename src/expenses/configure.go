package expenses

import (
	"database/sql"
	"app"
)

var db *sql.DB

func Configure(env *app.Env) {
	db = env.GetDB()

	routes := env.GetRoutes()
	routes.Handle("/expenses", Index(env)).Methods("GET")
	routes.Handle("/expenses", Create(env)).Methods("POST")
}
