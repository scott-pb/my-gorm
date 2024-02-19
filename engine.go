package mygorm

import (
	"database/sql"
	"github.com/scott-pb/mygorm/log"
	"github.com/scott-pb/mygorm/session"
)

// Engine 是数据库引擎的结构体
type Engine struct {
	db *sql.DB
}

// NewEngine 创建一个新的数据库引擎实例
func NewEngine(driver, source string) (*Engine, error) {
	// 打开数据库连接
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// 测试数据库连接是否可用
	if err = db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info("连接数据库成功")
	return &Engine{
		db: db,
	}, nil
}

// GetSession 返回一个数据库会话实例
func (e *Engine) GetSession() *session.Session {
	return session.New(e.db)
}

// Close 关闭数据库连接
func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("关闭数据库失败", err)
		return
	}

	log.Info("关闭数据库成功")
}
