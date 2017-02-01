package db

import (
	"database/sql"
	"dude_expenses/app"
	_ "github.com/lib/pq"
)

func Configure(env *app.Env) {
	db, err := sql.Open("postgres", "postgres://"+env.GetDBUrl()+"?sslmode=require")
	if err != nil {
		app.HandleFatal(err)
	}
	if err = db.Ping(); err != nil {
		Close(db)
		app.HandleFatal(err)
	}

	env.SetDB(db)
}

func Close(db *sql.DB) {
	db.Close()
}
