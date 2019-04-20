package main

import (
	"fmt"
	"github.com/XiaoXice/AirNet/common/util"
	"testing"
)

func TestValue(t *testing.T) {
	type a struct {
		A string
		b string
	}
	node := a{"1","2"}
	fmt.Println(util.Strict2Map(node))
	//v := reflect.ValueOf(node)
	//T := reflect.TypeOf(node)
	//for count := v.NumField() - 1; count >= 0; count -- {
	//	field := v.Field(count)
	//	newType := reflect.New(field.Type())
	//	fmt.Println(T.Field(count).Name,"+", field.Interface(), "+", reflect.TypeOf(newType.Interface()))
	//}
}

