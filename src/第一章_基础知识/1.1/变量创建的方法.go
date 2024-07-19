package main

import "fmt"

func main() {
	fun6()
}

// 1. var 是关键字（固定不变），name 是变量名，type 是类型。
func fun1() {
	var name string = "Go编程时光"
	fmt.Println(name)

	// 可以将其简化为
	var name2 = "Go编程时光"
	fmt.Println(name2)

	//若你的右值带有小数点，在不指定类型的情况下，编译器会将你的这个变量声明为 float64，但是很多情况下，我们并不需要这么高的精度（占用的内存空间更大）
	var num float32 = 0.11
	fmt.Println(num)

}

// 2多个变量一起声明
// func fun2() {
// 	var (
// 		name   string
// 		age    int
// 		gender string
// 	)
// }

// 3 去掉var 使用 := 声明和初始化一个变量
func fun3() {
	name3 := "3113"
	fmt.Println(name3)
}

// 4 去掉var 使用 := 声明和初始化多个变量
// 交换值的用法
func fun4() {
	// name, age := "wangbm", 28
	// fmt.Println(name)
	// fmt.Println(age)

	var a int = 100
	var b int = 200
	fmt.Print("交换前", a, b)
	b, a = a, b
	fmt.Print("交换后", a, b)

}

// 5 new 函数声明一个指针变量
// 变量分为两种 普通变量 和 指针变量
// & 和 * 的用法
// 变量/常量都只能声明一次，声明多次，编译就会报错。
func fun5() {
	// var age int = 28
	// var ptr = &age // &后面接变量名，表示取出该变量的内存地址
	// fmt.Println("age: ", age)
	// fmt.Println("ptr: ", ptr)

	//  new 函数，是 Go 里的一个内建函数。
	ptr := new(int)
	fmt.Println("ptr address: ", ptr)
	fmt.Println("ptr value: ", *ptr) // *后面接指针变量，表示从内存地址中取出值
}

//	匿名变量  也称作占位符，或者空白标识符，用下划线表示。
//
// 优点有三：
// 不分配内存，不占用内存空间
// 不需要你为命名无用的变量名而纠结
// 多次声明不会有任何问题
// 通常我们用匿名接收必须接收，但是又不会用到的值。
func fun6() {
	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)
}
func GetData() (int, int) {
	return 100, 200
}

// new：为所有的类型分配内存，并初始化为零值，返回指针。
// make：只能为 slice，map，chan 分配内存，并初始化，返回的是类型。
type Student struct {
	name string
	age  int
}

func fun7() {

	// new 一个内建类型
	num := new(int)
	// 等价于
	// num := 1
	fmt.Println(*num) //打印零值：0

	// new 一个自定义类型
	s := new(Student)
	s.name = "wangbm"
	// 切片
	a := make([]int, 2, 10)
	fmt.Println(a)
	// 字典
	b := make(map[string]int)
	fmt.Println(b)
	// 通道
	c := make(chan int, 10)
	fmt.Println(c)
}
