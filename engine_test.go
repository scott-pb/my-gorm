package mygorm

import (
	"github.com/scott-pb/mygorm/log"
	"testing"
)

func TestEngine(t *testing.T) {
	// 创建一个新的数据库引擎实例
	engine, _ := NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/my-gorm")
	// 延迟关闭数据库引擎连接
	defer engine.Close()

	// 获取一个数据库会话
	s := engine.GetSession()
	// 执行一条 SQL 查询，并获取结果的单行数据
	result := s.Raw("select * from mg_user").QueryRow()

	// 定义一个 User 结构体，用于存储从数据库查询结果中获取的用户信息
	type User struct {
		id       uint   // 用户 ID
		name     string // 用户名
		birthday string // 出生日期
		gender   uint8  // 性别
	}
	var user User

	// 将查询结果中的字段值扫描并赋值给 User 结构体中的字段
	if err := result.Scan(&user.id, &user.name, &user.birthday, &user.gender); err != nil {
		// 如果出现错误，则记录错误信息并返回
		log.Error(err)
		return
	}
	// 输出查询到的用户信息
	log.Info(user)
}
