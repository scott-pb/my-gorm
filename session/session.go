package session

import (
	"database/sql"
	"github.com/scott-pb/mygorm/log"
	"strings"
)

// Session 包含一个数据库连接（db）、SQL语句构建器（sql）和SQL变量列表（vars）。
type Session struct {
	db   *sql.DB
	sql  strings.Builder
	vars []interface{}
}

// New 创建并返回一个新的 Session 实例，接受一个数据库连接作为参数。
func New(db *sql.DB) *Session {
	return &Session{db: db}
}

// Reset 重置会话，清空之前的SQL语句和变量。
func (s *Session) Reset() {
	s.sql.Reset()
	s.vars = nil
}

// GetDB 返回当前会话的数据库连接。
func (s *Session) GetDB() *sql.DB {
	return s.db
}

// Raw 一个 SQL 查询字符串和可变数量的参数，并将它们添加到 Session 结构体的内部状态
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.vars = append(s.vars, values...)
	return s
}

// Exec 执行数据库写操作，并返回 sql.Result 对象和可能的错误
func (s *Session) Exec() (sql.Result, error) {
	// 在函数返回前调用 Reset 方法，确保资源清理
	defer s.Reset()

	// 调用 GetDB 方法获取数据库连接，执行 SQL 语句并返回结果
	r, err := s.GetDB().Exec(s.sql.String(), s.vars...)
	if err != nil {
		// 如果发生错误，记录错误信息
		log.Error(err)
	}

	// 记录调试信息，包括执行的 SQL 语句和参数
	log.Debug(s.sql.String(), s.vars)
	return r, err
}

// Query 执行数据库查询操作，并返回 *sql.Rows 对象和可能的错误
func (s *Session) Query() (*sql.Rows, error) {
	// 在函数返回前调用 Reset 方法，确保资源清理
	defer s.Reset()

	// 调用 GetDB 方法获取数据库连接，执行 SQL 语句并返回结果集
	rows, err := s.GetDB().Query(s.sql.String(), s.vars...)
	if err != nil {
		// 如果发生错误，记录错误信息
		log.Error(err)
	}

	return rows, err
}

// QueryRow 执行数据库查询操作，并返回 *sql.Row 对象
func (s *Session) QueryRow() *sql.Row {
	// 在函数返回前调用 Reset 方法，确保资源清理
	defer s.Reset()

	// 调用 GetDB 方法获取数据库连接，执行 SQL 语句并返回单行结果
	row := s.GetDB().QueryRow(s.sql.String(), s.vars...)

	// 记录调试信息，包括执行的 SQL 语句和参数
	log.Debug(s.sql.String(), s.vars)
	return row
}
