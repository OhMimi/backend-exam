package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

/*
請完成 TrimAllStrings 函式, 移除所有空白字元

禁止修改 TrimAllStrings 類型標記
*/

func TrimAllStrings(a any, visited map[uintptr]struct{}) {
	// 呼叫遞迴輔助函式
	trim(reflect.ValueOf(a), visited)
}

// 遞迴輔助函式
func trim(v reflect.Value, visited map[uintptr]struct{}) {
	// 持續解引用指標，直到 v 不是指標為止
	for v.Kind() == reflect.Ptr {
		// 如果 v 是 nil 指標，則直接返回
		if v.IsNil() {
			return
		}
		// 取得指標的地址
		addr := v.Pointer()
		// 檢查此指標地址是否已經訪問過
		if _, ok := visited[addr]; ok {
			return
		}
		// 標記此指標地址為已訪問
		visited[addr] = struct{}{}
		// 取得指標指向的元素
		v = v.Elem()
	}

	// 根據值的類型進行處理
	switch v.Kind() {
	case reflect.Struct:
		// 遍歷其所有欄位
		for i := 0; i < v.NumField(); i++ {
			trim(v.Field(i), visited)
		}
	case reflect.String:
		// 檢查該字串是否可以被修改
		if v.CanSet() {
			v.SetString(strings.TrimSpace(v.String()))
		}
	case reflect.Slice, reflect.Array:
		// 如果是切片或陣列，遍歷其所有元素
		for i := 0; i < v.Len(); i++ {
			trim(v.Index(i), visited)
		}
	case reflect.Map:
		// 如果是映射，遍歷其所有鍵值對
		for _, key := range v.MapKeys() {
			trim(v.MapIndex(key), visited)
		}
	default:
	}
}

func main() {
	type Person struct {
		Name string
		Age  int
		Next *Person
	}

	a := &Person{
		Name: " name ",
		Age:  20,
		Next: &Person{
			Name: " name2 ",
			Age:  21,
			Next: &Person{
				Name: " name3 ",
				Age:  22,
			},
		},
	}

	visited := make(map[uintptr]struct{})

	TrimAllStrings(&a, visited)

	m, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	// 預期輸出: {"Name":"name","Age":20,"Next":{"Name":"name2","Age":21,"Next":{"Name":"name3","Age":22,"Next":null}}}
	fmt.Println(string(m))

	a.Next = a

	TrimAllStrings(&a, visited)

	// 預期輸出: true
	fmt.Println(a.Next.Next.Name == "name")
}
