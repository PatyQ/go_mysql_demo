package dao

import (
	"encoding/json"
	"fmt"
	"go_mysql_demo/utilinit"
)

func RedisExample() {

	defer utilinit.RDB.Close()
	//set := utilinit.RDB.Set("name", "小明同学", 0)
	//fmt.Println("RedisExample redis set:::", utilinit.JsonString(set))
	//key := utilinit.RDB.RandomKey()
	//fmt.Println("RedisExample RandomKey:::", utilinit.JsonString(key))
	//get := utilinit.RDB.Get("name")
	//fmt.Println("redis get:::", get)

	userMap := make(map[string]interface{})

	//userMap["612XIAOMING"] = User{
	//	Id:   1,
	//	Age:  20,
	//	Name: "小明",
	//}
	//userMap["613xiaohua"] = User{
	//	Id:   1,
	//	Age:  22,
	//	Name: "小华",
	//}
	//set := utilinit.RDB.HSet("user", "usermap",utilinit.JsonString(userMap))
	//fmt.Println("HMSet:::",set)

	hmGet := utilinit.RDB.HMGet("user", "usermap")
	result, err := hmGet.Result()
	if err != nil {
		fmt.Println("hmGet.Result() error", err)
		return
	}

	//marshal, _ := json.Marshal(result[0])
	err = json.Unmarshal([]byte(result[0].(string)), &userMap) //success
	if err != nil {
		fmt.Println("json.Unmarshal error", err)
	}

	//rmap,ok := result[0].(map[string]interface{})
	//fmt.Println("hmGet.Result()",result,userMap)

	//printJson := utilinit.GetPrintJson(userMap)
	//fmt.Println(printJson) //{"612XIAOMING":{"id":1,"age":20,"name":"小明"},"613xiaohua":{"id":1,"age":22,"name":"小华"}}
}
