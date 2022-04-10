package utilinit

import (
	"fmt"
	"github.com/spf13/viper"
	"go_mysql_demo/dao"
)

func ConfInit() {
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

}

func InitSQLDb() {
	err := dao.InitDb(viper.GetString("dao.url"))
	if err != nil {
		fmt.Println("err::", err)
		return
	}
}

func InitXSQLDb() {
	err := dao.InitXDB(viper.GetString("dao.url"))
	if err != nil {
		fmt.Println("err::", err)
		return
	}
}
