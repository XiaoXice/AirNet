package util

import (
	"bytes"
	"math/rand"
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

func RandomInList(l []float32, num *int) []int {
	if num == nil {
		var total float32 = 0
		for _, v := range l {
			total += v
		}
		random := rand.Float32() * total
		for index, v := range l {
			random -= v
			if random < 0 {
				return []int{index}
			}
		}
		return []int{len(l) - 1}
	} else if len(l) < *num {
		var list []int
		for index := 0; index < len(l); index += 1 {
			list = append(list, index)
		}
		return list
	} else {
		var list []int
		for c := 1; c < *num; c += 1 {
			var total float32 = 0
			for _, v := range l {
				total += v
			}
			random := rand.Float32() * total
			for index, v := range l {
				random -= v
				if random < 0 {
					l[index] = 0
					list = append(list, index)
				}
			}
		}
		return list
	}
}

func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

func Find(l []interface{}, o interface{}) int {
	for index, v := range l {
		switch o.(type) {
		case []byte:
			if bytes.Equal(v.([]byte),o.([]byte)) {
				return index
			}
		default:
			if o == v {
				return index
			}
		}
	}
	return -1
}


func DeleteSlice(slice interface{}, index int) interface{} {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	if slice == nil || length == 0 || (length-1) < index {
		return slice
	}
	if length-1 == index {
		return sliceValue.Slice(0, index).Interface()
	} else if (length - 1) >= index {
		return reflect.AppendSlice(sliceValue.Slice(0, index), sliceValue.Slice(index+1, length)).Interface()
	}
	return slice
}