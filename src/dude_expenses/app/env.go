package app

import (
	"database/sql"
	"github.com/gorilla/mux"
	"io"
	"os"
)

type Env struct {
	db        *sql.DB
	router    *mux.Router
	userId    string
	logStream io.Writer
}

func New() *Env {
	return &Env{router: mux.NewRouter(), logStream: os.Stdout}
}

func (env *Env) SetDB(db *sql.DB) {
	env.db = db
}

func (env *Env) GetDB() *sql.DB {
	return env.db
}

func (env *Env) GetRouter() *mux.Router {
	return env.router
}

func (env *Env) SetUserId(userId string) {
	env.userId = userId
}

func (env *Env) GetUserId() string {
	return env.userId
}

func (env *Env) GetLogStream() io.Writer {
	return env.logStream
}
