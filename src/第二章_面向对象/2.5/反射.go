package main

import (
	"fmt"
	"reflect"
)

// 在学习 reflect 包的使用时，先得学习下这两种类型：
// reflect.Type
// reflect.Value

func main() {
	fun5()
}

func fun1() {
	var age interface{} = 25

	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)
	fmt.Println("type:", t)
	fmt.Println("type:", v)

	fmt.Println("可写性为:", v.CanSet())

	fmt.Printf("原始接口变量的类型为 %T，值为 %v \n", age, age)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", t)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v)

	// 从反射对象到接口变量
	i := v.Interface()
	fmt.Printf("从反射对象到接口变量：新对象的类型为 %T 值为 %v \n", i, i)

}

// 要让反射对象具备可写性，需要注意两点
// 创建反射对象时传入变量的指针
// 使用 Elem()函数返回指针指向的数据
func fun2() {
	var name string = "Go编程时光"
	v1 := reflect.ValueOf(&name)
	fmt.Println("v1 可写性为:", v1.CanSet())

	v2 := v1.Elem()
	fmt.Println("v2 可写性为:", v2.CanSet())

	// v2.set  可以通过编译器提示 修改哪些
}

type Profile struct {
}
type Person struct {
	name   string
	age    int
	gender string
}

func (p Person) SayBye() string {
	fmt.Println("Bye")
	return "xx"
}

func (p Person) SayHello() string {
	fmt.Println("Hello")
	return "ss"
}

// 学习反射的函数
func fun3() {
	m := Profile{}
	// 1.获取类别：Kind()
	// Type 对象 和 Value 对象都可以通过 Kind() 方法返回对应的接口变量的基础类型。
	// Kind 和 Type 是有区别的，Kind 表示更基础，范围更广的分类。
	// Kind 表示的基本都是 Go 原生的基本类型（共有 25 种的合法类型），而不包含自定义类型。
	t := reflect.TypeOf(m)
	fmt.Println("Type: ", t)
	fmt.Println("Kind: ", t.Kind())

	// 2.传入指针，转换值 关于 Elem() 已经讲过了，它会返回 Type 对象所表示的指针指向的数据。

	// Int() ：转成 int

	var age int = 25
	v1 := reflect.ValueOf(age)
	fmt.Printf("转换前， type: %T, value: %v \n", v1, v1)
	v2 := v1.Int()
	fmt.Printf("转换后， type: %T, value: %v \n", v2, v2)
	// Float()：转成 float
	var score float64 = 99.5
	v3 := reflect.ValueOf(score)
	fmt.Printf("转换前， type: %T, value: %v \n", v3, v3)
	v4 := v3.Float()
	fmt.Printf("转换后， type: %T, value: %v \n", v4, v4)

	// String()：转成 string
	var name string = "Go编程时光"
	v5 := reflect.ValueOf(name)
	fmt.Printf("转换前， type: %T, value: %v \n", v5, v5)
	v6 := v5.String()
	fmt.Printf("转换后， type: %T, value: %v \n", v6, v6)
	// Bool()：转成布尔值
	// Pointer()：转成指针
	// Interface()：转成接口类型

	// 3. 对切片的操作
	// Slice()：对切片再切片（两下标）
	// Slice() 函数与上面所有类型转换的函数都不一样，它返回还是 reflect.Value 反射对象，而不再是我们所想的真实世界里的切片对象。

	var numList []int = []int{1, 2}

	v7 := reflect.ValueOf(numList)
	fmt.Printf("Slice转换前， type: %T, value: %v \n", v7, v7)

	// Slice 函数接收两个参数
	v8 := v7.Slice(0, 2)
	fmt.Printf("Slice转换后， type: %T, value: %v \n", v8, v8)
	// Slice3()：对切片再切片（三下标）
	// Set() 和 Append()：更新切片

	// 4. 对属性的操作
	// NumField() 和 Field()
	p := Person{"写代码的", 27, "male"}

	v := reflect.ValueOf(p)

	fmt.Println("字段数:", v.NumField())
	fmt.Println("第 1 个字段：", v.Field(0))
	fmt.Println("第 2 个字段：", v.Field(1))
	fmt.Println("第 3 个字段：", v.Field(2))

	fmt.Println("==========================")
	// 也可以这样来遍历
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("第 %d 个字段：%v \n", i+1, v.Field(i))
	}
	// 5. 对方法的操作
	// NumMethod() 和 Method()
	p2 := &Person{"写代码的", 27, "male"}
	t2 := reflect.TypeOf(p2)
	fmt.Println("方法数（可导出的）:", t2.NumMethod())
	fmt.Println("第 1 个方法：", t2.Method(0).Name)
	fmt.Println("第 2 个方法：", t2.Method(1).Name)

	fmt.Println("==========================")
	// 也可以这样来遍历
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("第 %d 个方法：%v \n", i+1, t.Method(i).Name)
	}
}

// 动态调用函数
func fun4() {
	p := &Person{"wangbm", 27, "male"}

	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	for i := 0; i < v.NumMethod(); i++ {
		fmt.Printf("调用第 %d 个方法：%v ，调用结果：%v\n",
			i+1,
			t.Method(i).Name,
			v.Elem().Method(i).Call(nil))
	}
}
func (p Person) SelfIntroduction(name string, age int) {
	fmt.Printf("Hello, my name is %s and i'm %d years old.", name, age)
}

// 动态调用函数（使用函数且有参数
func fun5() {
	p := &Person{}

	v := reflect.ValueOf(p)
	//t := reflect.TypeOf(p)
	name := reflect.ValueOf("wangbm")
	age := reflect.ValueOf(27)
	input := []reflect.Value{name, age}
	v.MethodByName("SelfIntroduction").Call(input)
}
