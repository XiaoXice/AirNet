package util

import (
	"reflect"
)

func Strict2Map(obj interface{}) (map[string]interface{}, []string) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	var list []string
	for i := 0; i < t.NumField(); i++ {
		n := t.Field(i).Name
		if int(n[0]) < int("a"[0]) {
			data[t.Field(i).Name] = v.Field(i).Interface()
			list = append(list, n)
		}
	}
	return data, list
}

func Map2Array(obj map[string]interface{}, l []string) []interface{} {
	var data []interface{}
	for _, v := range l {
		data = append(data, obj[v])
	}
	return data
}
