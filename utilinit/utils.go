package utilinit

import (
	"encoding/json"
	"fmt"
)

func JsonString(b interface{}) string {

	marshal, err := json.Marshal(b)
	if err != nil {
		fmt.Println("JsonString error:::", err)
		return ""
	}
	return string(marshal)
}

// GetPrintJson 获取json格式
func GetPrintJson(v interface{}) string {
	msg, _ := json.Marshal(v)
	return string(msg)
}
