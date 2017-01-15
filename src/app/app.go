package app

import (
	"database/sql"
	"github.com/gorilla/mux"
)

type Env struct {
	db     *sql.DB
	routes *mux.Router
}

func New() *Env {
	return &Env{}
}

func (env *Env) SetDB(db *sql.DB) {
	env.db = db
}

func (env *Env) GetDB() *sql.DB {
	return env.db
}

func (env *Env) SetRoutes(routes *mux.Router) {
	env.routes = routes
}

func (env *Env) GetRoutes() *mux.Router {
	return env.routes
}
