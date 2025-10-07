package main

import (
	"fmt"
	"reflect"
)

/*
請完成 swap 函式, 交換兩個變數的值

允許panic但必須是顯式調用
地址不允許改變
禁止修改 swap 類型標記
*/

func swap[T any](a, b T) {
	// 獲取 a 和 b 的反射值
	valA := reflect.ValueOf(a)
	valB := reflect.ValueOf(b)

	// 檢查傳入的是否為指標，如果不是，則無法進行修改，直接 panic
	if valA.Kind() != reflect.Ptr || valB.Kind() != reflect.Ptr {
		panic("swap: arguments must be pointers")
	}

	// 使用 .Elem() 方法來獲取 a 和 b 變數本身的值
	elemA := valA.Elem()
	elemB := valB.Elem()

	// 建立臨時變數來儲存A的值
	tempA := elemA.Interface()

	// 執行交換
	elemA.Set(elemB)
	elemB.Set(reflect.ValueOf(tempA))
}

func main() {
	a := 10
	b := 20

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)

	swap(&a, &b)

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)
}
