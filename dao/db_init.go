package dao

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB // 定义一个全局对象db

func InitDb(dburl string) error {
	var err error
	fmt.Println(dburl)
	// DSN:Data Source Name
	dsn := dburl
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = Db.Ping()
	if err != nil {
		return err
	}
	return nil
}
