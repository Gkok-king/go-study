package main

import "fmt"

func main() {
	// fun9()
	x := 1
	c := &x
	mytest(c)
}

// 整型 int（有符号） 和 uint
// 根据有无符号（正负） 和大小（8、16、32、64、32或64） 分为10种
// 而 int 并没有指定它的位数，说明它的大小，是可以变化的，那根据什么变化呢？
// 当你在32位的系统下，int 和 uint 都占用 4个字节，也就是32位。
// 若你在64位的系统下，int 和 uint 都占用 8个字节，也就是64位。
// 出于这个原因，在某些场景下，你应当避免使用 int 和 uint ，
// 而使用更加精确的 int32 和 int64，比如在二进制传输、读写文件的结构描述（为了保持文件的结构不会受到不同编译目标平台字节长度的影响）
func fun1() {
	// var num int = 10
	// fmt.Print(num)
	//可以用其他进制来表达
	//2进制：以0b或0B为前缀
	//8进制：以0o或者 0O为前缀
	//16进制：以0x 为前缀
	var num01 int = 0b1100
	var num02 int = 0o14
	var num03 int = 0xC
	fmt.Printf("2进制数 %b 表示的是: %d \n", num01, num01)
	fmt.Printf("8进制数 %o 表示的是: %d \n", num02, num02)
	fmt.Printf("16进制数 %X 表示的是: %d \n", num03, num03)

	// %b    表示为二进制
	// %c    该值对应的unicode码值
	// %d    表示为十进制
	// %o    表示为八进制
	// %q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	// %x    表示为十六进制，使用a-f
	// %X    表示为十六进制，使用A-F
	// %U    表示为Unicode格式：U+1234，等价于"U+%04X"
	// %E    用科学计数法表示
	// %f    用浮点数表示
}

// 浮点型
// Go语言中提供了两种精度的浮点数 float32 和 float64。
func fun2() {

	var myfloat01 float32 = 100000182
	var myfloat02 float32 = 100000187
	fmt.Println("myfloat: ", myfloat01)
	fmt.Println("myfloat: ", myfloat01+5)
	//由于其类型是 float32，精度不足，导致最后比较的结果是不相等（从小数点后第七位开始不精确）
	fmt.Println(myfloat02 == myfloat01+5)
}

// byte与rune.go
func fun3() {
	var a byte = 65
	// 8进制写法: var a byte = '\101'     其中 \ 是固定前缀
	// 16进制写法: var a byte = '\x41'    其中 \x 是固定前缀

	var b uint8 = 66
	fmt.Printf("a 的值: %c \nb 的值: %c", a, b)

	// 或者使用 string 函数
	// fmt.Println("a 的值: ", string(a)," \nb 的值: ", string(b))

	// byte 和 uint8 没有区别，rune 和 int32 没有区别，那为什么还要多出一个 byte 和 rune 类型呢？
	// 理由很简单，因为uint8 和 int32 ，直观上让人以为这是一个数值，但是实际上，它也可以表示一个字符，
}

// 字符型
func fun4() {
	var country string = "hello,中国"
	fmt.Println(len(country))
	//除了双引号之外 ，你还可以使用反引号。
	var mystr01 string = "\\r\\n"
	var mystr02 string = `\r\n`
	fmt.Println(mystr01)
	fmt.Println(mystr02)
	//解释型的字符串，但是各种转义实在太麻烦了。你可以使用 fmt 的 %q 来还原
	fmt.Printf("的解释型字符串是： %q", mystr01)
	//同时反引号可以不写换行符（因为没法写）来表示一个多行的字符串。
	var mystr03 string = `你好呀!
	我的公众号是: Go编程时光，欢迎大家关注`
	fmt.Println(mystr03)

}

// 数组
func fun5() {

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	// 声明的第一种方法
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr1)

	// 声明的第二种方法
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// 3 表示数组的元素个数 ，万一你哪天想往该数组中增加元素，你得对应修改这个数字，
	// 为了避免这种硬编码，你可以这样写，使用 ... 让Go语言自己根据实际情况来分配空间
	arr4 := [...]int{1, 2, 3}
	fmt.Println(arr4)

	// 使用 type 关键字可以定义一个类型字面量，后面只要你想定义一个容器大小为3，元素类型为int的数组 ，都可以使用这个别名类型。
	type arr3 [3]int

	myarr := arr3{1, 2, 3}
	fmt.Printf("%d 的类型是: %T", myarr, myarr)

	// 定义数组还有一种偷懒的方法
	// 4表示数组有4个元素，2 和 3 分别表示该数组索引为2（初始索引为0）的值为3，而其他没有指定值的，就是 int 类型的零值，即0
	arr5 := [4]int{2: 3}
	fmt.Printf("数组值", arr5)
}

// 切片（slice
// 与数组一样，也是可以容纳若干类型相同的元素的容器。与数组不同的是，无法通过切片类型来确定其值的长度。
// 每个切片值都会将数组作为其底层数据结构。我们也把这样的数组称为切片的底层数组）

// 切片是对数组的一个连续片段的引用，所以切片是一个引用类型，这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集，
// 需要注意的是，终止索引标识的项不包括在切片内（意思是这是个左闭右开的区间）
func fun6() {
	// // myarr := [...]int{1, 2, 3}
	// // fmt.Printf("%d 的类型是: %T", myarr[0:2], myarr[0:2])

	// //对数组进行片段截取，主要有如下两种写法
	// // 定义一个数组
	// myarr1 := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("创建", myarr1)
	// // 【第一种】
	// // 1 表示从索引1开始，直到到索引为 2 (3-1)的元素
	// mysli2 := myarr1[1:3]
	// fmt.Println("第2个切", mysli2)
	// // 【第二种】
	// // 1 表示从索引1开始，直到到索引为 2 (3-1)的元素
	// mysli3 := myarr1[1:3:4]
	// fmt.Println("第3个切", mysli3)
	// //在切片时，若不指定第三个数，那么切片终止索引会一直到原数组的最后一个数。而如果指定了第三个数，那么切片终止索引只会到原数组的该索引值。

	// myarr := [5]int{1, 2, 3, 4, 5}
	// fmt.Printf("myarr 的长度为：%d，容量为：%d\n", len(myarr), cap(myarr))

	// mysli1 := myarr[1:3]
	// fmt.Printf("mysli1 的长度为：%d，容量为：%d\n", len(mysli1), cap(mysli1))
	// fmt.Println(mysli1)

	// mysli2 := myarr[1:3:4]
	// fmt.Printf("mysli2 的长度为：%d，容量为：%d\n", len(mysli2), cap(mysli2))
	// fmt.Println(mysli2)

	// 切片的从头声明赋值
	// 声明字符串切片
	var strList []string
	fmt.Println(strList)
	// 声明整型切片
	var numList []int
	fmt.Println(numList)
	// 声明一个空切片
	var numListEmpty = []int{}
	fmt.Println(numListEmpty)

	// 使用 make 函数构造，make 函数的格式：make( []Type, size, cap )
	// 	关于 len 和 cap 的概念，可能不好理解 ，这里举个例子：
	// 公司名，相当于字面量，也就是变量名。
	// 公司里的所有工位，相当于已分配到的内存空间
	// 公司里的员工，相当于元素。
	// cap 代表你这个公司最多可以容纳多少员工
	// len 代表你这个公司当前有多少个员工
	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
	fmt.Println(cap(a), cap(b))

	// 使用和数组一样，偷懒的方法
	// slice1 := []int{4: 2}
	// arr5 := [4]int{2: 3}
	// fmt.Printf("切片值slice1", slice1)
	// fmt.Printf("数组值", arr5)

	//  创建一个数组
	myArr := [...]int{1, 2, 3}
	// 将数组转换成slice
	slice := myArr[:]

	fmt.Println("切片值slice1", slice)
	// 手动创建一个slice
	slice2 := []int{1, 23, 4}
	fmt.Println("切片值slice2", slice2)
}

// 字典(map)
func fun7() {
	// // 三种声明并初始化字典的方法
	// // 第一种方法
	// var scores map[string]int = map[string]int{"english": 80, "chinese": 85}
	// fmt.Println(scores)
	// // 第二种方法
	// scores2 := map[string]int{"english": 80, "chinese": 85}
	// fmt.Println(scores2)
	// // 第三种方法
	// scores3 := make(map[string]int)
	// scores3["english"] = 80
	// scores3["chinese"] = 85

	//crud 存在就更新
	scores4 := make(map[string]int)
	scores4["math"] = 95
	scores4["math"] = 100
	fmt.Println(scores4["math"])
	delete(scores4, "ss")

	//判断 key 是否存在
	// 实字典的下标读取可以返回两个值，使用第二个返回值都表示对应的 key 是否存在，若存在ok为true，若不存在，则ok为false
	if math, ok := scores4["math"]; ok {
		fmt.Printf("math 的值是: %d", math)
	} else {
		fmt.Println("math 不存在")
	}

	//获取 key 和 value
	scores5 := map[string]int{"english": 80, "chinese": 85}
	for subject, score := range scores5 {
		fmt.Printf("key: %s, value: %d\n", subject, score)
	}
	//只获取key，这里注意不用占用符。
	scores6 := map[string]int{"english": 80, "chinese": 85}
	for subject := range scores6 {
		fmt.Printf("key: %s\n", subject)
	}
	//只获取 value，用一个占位符替代。
	scores7 := map[string]int{"english": 80, "chinese": 85}
	for _, score := range scores7 {
		fmt.Printf("value: %d\n", score)
	}
}

// 布尔类型
// 在 Go 中，真值用 true 表示，不但不与 1 相等，并且更加严格，不同类型无法进行比较，而假值用 false 表示，同样与 0 无法比较。
// bool 与 int 不能直接转换，如果要转换，需要你自己实现函数。
func fun8() {
	//逻辑值取反 用！
	fmt.Println(true != false)
	//使用 && 表示且，用 || 表示或，并且有短路行为（即左边表达式已经可以确认整个表达式的值，那么右边将不会再被求值。
	//  && 两边的表达式都会执行
	age := 55
	gender := "male"
	fmt.Println(age > 18 && gender == "male")
	// gender == "male" 并不会执行
	fmt.Println(age < 18 || gender == "male")
}

// 指针
// 指针变量（一个标签）的值是指针，也就是内存地址
func fun9() {
	//指针的创建
	// 1先定义对应的变量，再通过变量取得内存地址，创建指针
	// 定义普通变量
	aint := 1
	// 定义指针变量
	ptr := &aint
	fmt.Println(ptr)

	// 2先创建指针，分配好内存后，再给指针指向的内存地址写入对应的值。
	// 创建指针
	astr := new(string)
	// 给指针赋值
	*astr = "Go编程时光"
	fmt.Println(astr)

	// 3先声明一个指针变量，再从其他变量取得内存地址赋值给它
	aint2 := 1
	var bint2 *int // 声明一个指针
	bint2 = &aint2 // 初始化
	fmt.Println(bint2)

	//	&：从一个普通变量中取得内存地址
	//  *：当*在赋值操作符（=）的右边，是从一个指针变量中取得变量值，当*在赋值操作符（=）的左边，是指该指针指向的变量

	// aint3 := 1     // 定义普通变量
	// ptr3 := &aint3 // 定义指针变量
	// fmt.Println("普通变量存储的是：", aint3)
	// fmt.Println("普通变量存储的是：", *ptr3)
	// fmt.Println("指针变量存储的是：", &aint3)
	// fmt.Println("指针变量存储的是：", ptr3)

	// 打印指针指向的内存地址，方法有两种
	// 第一种
	fmt.Printf("%p", ptr)
	// 第二种
	fmt.Println("这是指针", ptr)

	//指针的类型
	astr4 := "hello"
	aint4 := 1
	abool := false
	arune := 'a'
	afloat := 1.2

	fmt.Printf("astr 指针类型是：%T\n", &astr4)
	fmt.Printf("aint 指针类型是：%T\n", &aint4)
	fmt.Printf("abool 指针类型是：%T\n", &abool)
	fmt.Printf("arune 指针类型是：%T\n", &arune)
	fmt.Printf("afloat 指针类型是：%T\n", &afloat)

	//当指针声明后，没有进行初始化，其零值是 nil。
	a := 25
	var b *int // 声明一个指针
	if b == nil {
		fmt.Println(b)
		b = &a // 初始化：将a的内存地址给b
		fmt.Println(b)
	}
}

// 定义一个只接收指针类型的参数的函数，可以这么写
func mytest(ptr *int) {
	fmt.Println(*ptr)
}
