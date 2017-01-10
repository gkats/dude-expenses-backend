package app

import (
	"database/sql"
)

type Env struct {
	db *sql.DB
}

func New() *Env {
	return &Env{}
}

func (env *Env) SetDB(db *sql.DB) {
	env.db = db
}
