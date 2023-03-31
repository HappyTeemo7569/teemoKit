package convert

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func JSONToMap(content interface{}) map[string]interface{} {
	var name map[string]interface{}
	if marshalContent, err := json.Marshal(content); err != nil {
		fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber() // 设置将float64转为一个number
		if err := d.Decode(&name); err != nil {
			fmt.Println(err)
		} else {
			for k, v := range name {
				name[k] = v
			}
		}
	}
	return name
}

// JsonStr JSON化对象为字符串
func JsonStr(obj interface{}) string {
	b, _ := json.Marshal(obj)

	return string(b)
}
