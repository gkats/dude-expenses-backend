package db

import (
	"database/sql"
	"dude_expenses/app"
	_ "github.com/lib/pq"
)

func Configure(env *app.Env) {
	db, err := sql.Open("postgres", "postgres://"+env.GetDBUrl()+"?sslmode=require")
	if err != nil {
		// TODO Logging here...
	}
	if err = db.Ping(); err != nil {
		// TODO Logging here...
	}

	env.SetDB(db)
}

// TODO defer db.Close() in main... expose a Close() function for main to defer
