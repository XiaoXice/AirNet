package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/XiaoXice/AirNet/common/util"
	"github.com/XiaoXice/AirNet/net/database"
	"testing"
)

func TestValue(t *testing.T) {
	type a struct {
		A string
		b string
	}
	node := a{"1", "2"}
	fmt.Println(util.Strict2Map(node))
	//v := reflect.ValueOf(node)
	//T := reflect.TypeOf(node)
	//for count := v.NumField() - 1; count >= 0; count -- {
	//	field := v.Field(count)
	//	newType := reflect.New(field.Type())
	//	fmt.Println(T.Field(count).Name,"+", field.Interface(), "+", reflect.TypeOf(newType.Interface()))
	//}
}

func TestHashInfoEq(t *testing.T) {
	hash := make([]byte, 128)
	md := sha256.New()
	_, _ = rand.Reader.Read(hash)
	md.Write(hash)
	a := database.NewHashInfo(md.Sum(nil))
	b := database.NewHashInfo(md.Sum(nil))
	//ad := reflect.ValueOf(a).Interface()
	//bd := reflect.ValueOf(b).Interface()
	//b[31] = 110
	//var ap *database.HashInfo
	//var bp *database.HashInfo
	//ap = &a
	//bp = &b
	list := [][]byte{a.ToBytes(), b.ToBytes()}
	fmt.Println(util.Find(util.ToSlice(list), a.ToBytes()))
	//fmt.Println(bytes.Equal(ad,bd))
}
