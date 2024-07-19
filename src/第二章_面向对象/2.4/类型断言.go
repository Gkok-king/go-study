package main

import (
	"fmt"
)

// Type Assertion（中文名叫：类型断言），通过它可以做到以下几件事情
// 检查 i 是否为 nil
// 检查 i 存储的值是否为某个类型
func main() {
	fun3()
}

// 第一种使用：t := i.(T)
func fun1() {
	var i interface{} = 10
	t1 := i.(int)
	fmt.Println(t1)

	fmt.Println("=====分隔线=====")

	t2 := i.(string)
	fmt.Println(t2)
}

// 第二种使用：t, ok:= i.(T)
// 但和第一种表达式不同的事，这个不会触发 panic，而是将 ok 的值设为 false ，表示断言失败
func fun2() {
	var i interface{} = 10
	t1, ok := i.(int)
	fmt.Printf("%d-%t\n", t1, ok)

	fmt.Println("=====分隔线1=====")

	t2, ok := i.(string)
	fmt.Printf("%s-%t\n", t2, ok)

	fmt.Println("=====分隔线2=====")

	var k interface{} // nil
	t3, ok := k.(interface{})
	fmt.Println(t3, "-", ok)

	fmt.Println("=====分隔线3=====")
	k = 10
	t4, ok := k.(interface{})
	fmt.Printf("%d-%t\n", t4, ok)

	t5, ok := k.(int)
	fmt.Printf("%d-%t\n", t5, ok)
}

// Type Switch
// 区分多种类型，可以使用 type switch 断言，这个将会比一个一个进行类型断言更简单、直接、高效。
func fun3() {
	findType(10)      // int
	findType("hello") // string
	var k interface{} // nil
	findType(k)
	findType(10.23) //float64
}

func findType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case nil:
		fmt.Println(x, "is nil")
	default:
		fmt.Println(x, "not type matched")
	}
}
