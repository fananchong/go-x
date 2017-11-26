package util

import (
	"encoding/json"
	"reflect"
)

func Struct2String(obj interface{}) (string, error) {
	tmpdata := Struct2Map(obj)
	data, err := json.Marshal(tmpdata)
	if err == nil {
		return string(data), nil
	}
	return "", err
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
