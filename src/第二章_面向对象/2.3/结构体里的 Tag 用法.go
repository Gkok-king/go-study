package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	fun2()
}

type Person struct {
	Name string
	Age  int
	Addr string
}

// 1字段上还可以额外再加一个属性，用反引号（Esc键下面的那个键）包含的字符串，称之为 Tag，也就是标签。
type Person2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

// 由于 Person 结构体里的 Addr 字段有 omitempty 属性，因此 encoding/json 在将对象转化 json 字符串时，
// 只要发现对象里的 Addr 为 false， 0， 空指针，空接口，空数组，空切片，空映射，空字符串中的一种，就会被忽略。
func fun1() {
	p1 := Person2{
		Name: "Jack",
		Age:  22,
	}

	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	// p1 没有 Addr，就不会打印了
	fmt.Printf("%s\n", data1)

	// ================

	p2 := Person2{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	// p2 则会打印所有
	fmt.Printf("%s\n", data2)
}

// 2如何定义获取 Tag

func fun2() {
	p1 := Person2{
		Name: "Jack",
		Age:  22,
	}

	// 三种获取 field
	field, _ := reflect.TypeOf(p1).FieldByName("Name")
	// field := reflect.ValueOf(obj).Type().Field(i)         // i 表示第几个字段
	// field := reflect.ValueOf(obj).Elem().Type().Field(i) // i 表示第几个字段

	// 获取 Tag
	tag := field.Tag

	// 获取键值对
	labelValue := tag.Get("name")
	labelValue, ok := tag.Lookup("name")
	fmt.Println("labelValue", labelValue)
	fmt.Println("ok", ok)
}
