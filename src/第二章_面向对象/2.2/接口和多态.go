package main

import (
	"fmt"
	"strconv"
)

// 接口一般这样定义：接口定义一个对象的行为。接口只指定了对象应该做什么，至于如何实现这个行为（即实现细节），则由对象本身去确定。
func main() {
	// demo()
	fun1()
}

// 1. 如何定义接口
type Phone interface {
	call()
}

// 2. 实现接口
type Nokia struct {
	name string
}

// 接收者为 Nokia
// 接口的实现是隐式的，不像 JAVA 中要用 implements 显示说明。
func (phone Nokia) call() {
	fmt.Println("我是 Nokia，是一台电话")
}

// 3。接口实现多态

// 定义一个商品（Good）的接口
type Good interface {
	settleAccount() int
	orderInfo() string
}

// 定义两个结构体
type Phone2 struct {
	name     string
	quantity int
	price    int
}

type FreeGift struct {
	name     string
	quantity int
	price    int
}

// 然后分别为他们实现 Good 接口的两个方法
func (phone2 Phone2) settleAccount() int {
	return phone2.quantity * phone2.price
}
func (phone2 Phone2) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone2.quantity) + "个" +
		phone2.name + "计：" + strconv.Itoa(phone2.settleAccount()) + "元"
}

// FreeGift
func (gift FreeGift) settleAccount() int {
	return 100
}
func (gift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity) + "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

// 实例化
func demo() {
	iPhone2 := Phone2{
		name:     "iPhone",
		quantity: 1,
		price:    8000,
	}
	earphones := FreeGift{
		name:     "耳机",
		quantity: 1,
		price:    200,
	}
	//创建一个购物车（也就是类型为 Good的切片）
	goods := []Good{iPhone2, earphones}
	//调用计算金额的方法
	allPrice := calculateAllPrice(goods)
	fmt.Printf("该订单总共需要支付 %d 元", allPrice)
}

// 定义一个计算金额大方法
func calculateAllPrice(goods []Good) int {
	var allPrice int
	for _, good := range goods {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}

// 空接口
// 空接口是特殊形式的接口类型，普通的接口都有方法，而空接口没有定义任何方法口，也因此，我们可以说所有类型都至少实现了空接口。
func fun1() {
	var i interface{}
	fmt.Printf("type: %T, value: %v", i, i)
}

type empty_iface interface {
}

// 如何使用空接口
// 1通常我们会直接使用 interface{} 作为类型声明一个实例，而这个实例可以承载任意类型的值。
func myfunc2(iface interface{}) {
	fmt.Println(iface)
}

func fun2() {
	a := 10
	b := "hello"
	c := true

	myfunc2(a)
	myfunc2(b)
	myfunc2(c)
}

// 2如果想让你的函数可以接收任意类型的值 ，也可以使用空接口
func myfunc3(ifaces ...interface{}) {
	for _, iface := range ifaces {
		fmt.Println(iface)
	}
}
func fun3() {
	a := 10
	b := "hello"
	c := true
	myfunc3(a, b, c)

	// 3定义一个可以接收任意类型的 array、slice、map、strcut，例如这边定义一个切片
	any := make([]interface{}, 5)
	any[0] = 11
	any[1] = "hello world"
	any[2] = []int{11, 22, 33, 44}
	for _, value := range any {
		fmt.Println(value)
	}
}

// 空接口几个要注意的坑
// 坑1：空接口可以承载任意值，但不代表任意类型就可以承接空接口类型的值
// 坑2：：当空接口承载数组和切片后，该对象无法再进行切片
// 坑3：当你使用空接口来接收任意类型的参数时，它的静态类型是 interface{}，
// 但动态类型（是 int，string 还是其他类型）我们并不知道，因此需要使用类型断言。
func fun5() {
	// 声明a变量, 类型int, 初始值为1
	// var a int = 1
	// 声明i变量, 类型为interface{}, 初始值为a, 此时i的值变为1
	// var i interface{} = a
	// 声明b变量, 尝试赋值i
	// var b int = i

	// ===================
	// sli := []int{2, 3, 5, 7, 11, 13}
	// var i interface{}
	// i = sli
	// g := i[1:3]
	// fmt.Println(g)
}

// 接口的一些默认规则1 对方法的调用限制
type Phone5 interface {
	call()
}

type iPhone5 struct {
	name string
}

func (phone iPhone5) call() {
	fmt.Println("Hello, iPhone.")
}

func (phone iPhone5) send_wechat() {
	fmt.Println("Hello, Wechat.")
}

func fun6() {
	var phone Phone
	phone = iPhone5{name: "ming's iphone"}
	phone.call()
	// 调用 phone.send_wechat方法，程序会报错，提示我们 Phone5 类型的方法没有 send_wechat 的字段或方法。
	// 原因也很明显，因为我们的phone对象显式声明为 Phone5 接口类型，因此 phone5调用的方法会受到此接口的限制。
	// phone.send_wechat()

	phone2 := iPhone5{name: "ming's iphone"}
	phone2.call()
	phone2.send_wechat()
}

// 接口的一些默认规则2 调用函数时的隐式转换
// Go 语言中的函数调用都是值传递的，变量会在方法调用前进行类型转换。
func printType(i interface{}) {
	// 接收参数时会隐式转换
	switch i.(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}
func fun7() {
	a := 10
	printType(a)

	//这样会报错
	// b := 10
	// switch b.(type) {
	// case int:
	// 	fmt.Println("参数的类型是 int")
	// case string:
	// 	fmt.Println("参数的类型是 string")
	// }

	//需要做的显示转换
	b := 10
	switch interface{}(b).(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}
