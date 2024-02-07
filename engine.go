package mygorm

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/scott-pb/mygorm/log"
	"github.com/scott-pb/mygorm/session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (*Engine, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info("connect database success")
	return &Engine{
		db: db,
	}, nil
}

func (e *Engine) GetSession() *session.Session {
	return session.New(e.db)
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("close database failed", err)
		return
	}

	log.Info("close database success")

}
