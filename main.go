package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_mysql_demo/dao"
	"go_mysql_demo/utilinit"
)

func main() {
	utilinit.ConfInit()
	utilinit.InitSQLDb()
	utilinit.InitXSQLDb()
	utilinit.RedisInit()

	dao.RedisExample()
	//InsertSqlUser()
	//InsertXSqlUser()
}

func InsertSqlUser() {

	sqlStr := "insert into user(name,age)values(?,?)"
	exec, err := utilinit.Db.Exec(sqlStr, "王五就开了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊", 31)
	if err != nil {
		fmt.Println("InsertSqlUser error", "error", err)
		return
	}
	lastInsertId, err := exec.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId error", err)
		return
	}
	fmt.Println("LastInsertId:::", lastInsertId)
}

// 批量添加
func InsertXSqlUser() {
	//sqlStr := "insert into user(name,age)values(?,?)"
	//exec, err := dao.XDB.Exec(sqlStr, "小华", 18)
	users := make([]dao.User, 0)
	users = append(users, dao.User{
		Age:  10,
		Name: "小红",
	},
		dao.User{
			Age:  20,
			Name: "小绿",
		})
	exec, err := utilinit.XDB.NamedExec(`insert into user (name,age)values (:name,:age)`, users)
	if err != nil {
		fmt.Println("InsertXSqlUser error", "error", err)
		return
	}
	insertId, err := exec.LastInsertId()
	if err != nil {
		fmt.Println("XLastInsertId error", "error", err)
		return
	}
	fmt.Println("insertId", insertId)
}
