package db

import (
	"database/sql"
	"app"
	_"github.com/lib/pq"
)

func Configure(env *app.Env) {
	db, err := sql.Open("postgres", "postgres://dude_expenses:dudeExpen$es123@localhost:5432/dude_expenses_development?sslmode=require")
	if err != nil {
		// TODO Logging here...
	}
	if err = db.Ping(); err != nil {
		// TODO Logging here...
	}

	env.SetDB(db)
}

// TODO defer db.Close()