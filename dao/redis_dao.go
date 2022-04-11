package dao

import (
	"fmt"
	"go_mysql_demo/utilinit"
)

func RedisExample() {

	set := utilinit.RDB.Set("name", "小明同学", 0)
	fmt.Println("RedisExample redis set:::", utilinit.JsonString(set))
	key := utilinit.RDB.RandomKey()
	fmt.Println("RedisExample RandomKey:::", utilinit.JsonString(key))
	get := utilinit.RDB.Get("name")
	fmt.Println("redis get:::", get)
}
