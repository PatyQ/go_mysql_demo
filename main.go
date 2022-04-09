package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go_mysql_demo/dao"
)

func main() {
	viper.SetConfigFile("./mysqldemo.yaml") // 指定配置文件路径
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			fmt.Println("error:配置文件未找到错误；如果需要可以忽略", err)
		} else {
			// 配置文件被找到，但产生了另外的错误
			fmt.Println("error:配置文件被找到，但产生了另外的错误", err)
		}
	}
	err := dao.InitDb(viper.GetString("dao.url"))
	if err != nil {
		fmt.Println("err::", err)
		return
	}
	InsertSqlUser()

}

func InsertSqlUser() {

	sqlStr := "insert into user(name,age)values(?,?)"
	exec, err := dao.Db.Exec(sqlStr, "王五", 30)
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
