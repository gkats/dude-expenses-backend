package app

import (
	"database/sql"
	"github.com/gorilla/mux"
	"io"
	"os"
)

type Env struct {
	db         *sql.DB
	router     *mux.Router
	userId     string
	logStream  io.Writer
	authSecret string
	dbUrl      string
}

func New(authSecret string, dbUrl string) *Env {
	return &Env{
		router:     mux.NewRouter(),
		logStream:  os.Stdout,
		authSecret: authSecret,
		dbUrl:      dbUrl,
	}
}

func (env *Env) SetDB(db *sql.DB) {
	env.db = db
}

func (env *Env) GetDB() *sql.DB {
	return env.db
}

func (env *Env) GetDBUrl() string {
	return env.dbUrl
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

func (env *Env) GetAuthSecret() string {
	return env.authSecret
}
