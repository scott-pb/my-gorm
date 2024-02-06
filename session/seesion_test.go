package session

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/scott-pb/mygorm/log"
	"testing"
)

func TestSession_QueryRowQuery(t *testing.T) {
	// 打开 MySQL 数据库连接
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/my-gorm")
	if err != nil {
		log.Error(err)
		return
	}
	// 设置与数据库的最大打开连接数
	db.SetMaxOpenConns(100)
	// 设置空闲中的最大连接数
	db.SetMaxIdleConns(10)

	// 创建一个新的 Session 实例
	session := New(db)

	// 定义 User 结构体，用于扫描查询结果
	type User struct {
		id       uint   // 用户 ID
		name     string // 用户名
		birthday string // 出生日期
		gender   uint8  // 性别
	}

	var user User

	// 查询数据库表 mg_user 中的第一行数据
	row := session.Raw("select * from mg_user limit 1").QueryRow()

	// 扫描查询结果并将数据填充到 User 结构体中
	row.Scan(&user.id, &user.name, &user.birthday, &user.gender)

	// 输出查询到的用户信息
	log.Info(user)
}
