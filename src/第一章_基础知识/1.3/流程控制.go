package main

import (
	"fmt"
	"time"
)

func main() {
	fun7()
}

// if-else
// go对于 { 和 } 的位置有严格的要求，它要求 else if （或 else）和 两边的花括号，必须在同一行。
func fun1() {
	// if 条件 1 {
	// 	分支 1
	//   } else if 条件 2 {
	// 	分支 2
	//   } else if 条件 ... {
	// 	分支 ...
	//   } else {
	// 	分支 else
	//   }
	age := 20
	gender := "male"
	if age > 18 && gender == "male" {
		fmt.Println("是成年男性")
	}
	// if 里可以允许先运行一个表达式，取得变量后，再对其进行判断
	if age := 20; age > 18 {
		fmt.Println("已经成年了")
	}
}

// switch-case
func fun2() {

	// switch 表达式 {
	// case 表达式1:
	//     代码块
	// case 表达式2:
	//     代码块
	// case 表达式3:
	//     代码块
	// case 表达式4:
	//     代码块
	// case 表达式5:
	//     代码块
	// default:
	//     代码块
	// }

	education := "本科"

	switch education {
	case "博士":
		fmt.Println("我是博士")
	case "研究生":
		fmt.Println("我是研究生")
	case "本科":
		fmt.Println("我是本科生")
	case "大专":
		fmt.Println("我是大专生")
	case "高中":
		fmt.Println("我是高中生")
	default:
		fmt.Println("学历未达标..")
	}

	//case 后可以接多个多个条件，多个条件之间是 或 的关系，用逗号相隔。
	month := 2
	switch month {
	case 3, 4, 5:
		fmt.Println("春天")
	case 6, 7, 8:
		fmt.Println("夏天")
	case 9, 10, 11:
		fmt.Println("秋天")
	case 12, 1, 2:
		fmt.Println("冬天")
	default:
		fmt.Println("输入有误...")
	}

	// 当 case 后接的是常量时，该常量只能出现一次。
	// gender := "male"
	// switch gender {
	// 	case "male":
	// 		fmt.Println("男性")
	// 	// 与上面重复
	// 	case "male":
	// 		fmt.Println("男性")
	// 	case "female":
	// 		fmt.Println("女性")
	// }

	// switch 后可接函数
	chinese := 80
	english := 50
	math := 100
	switch getResult(chinese, english, math) {
	// case 后也必须 是布尔类型
	case true:
		fmt.Println("该同学所有成绩都合格")
	case false:
		fmt.Println("该同学有挂科记录")
	}

	//switch 后可以不接任何变量、表达式、函数。 当不接任何东西时，switch - case 就相当于 if - elseif - else
	score := 30
	switch {
	case score >= 95 && score <= 100:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("合格")
	case score >= 0:
		fmt.Println("不合格")
	default:
		fmt.Println("输入有误...")
	}

	// 正常情况下 switch - case 的执行顺序是：只要有一个 case 满足条件，就会直接退出 switch - case ，如果 一个都没有满足，才会执行 default 的代码块。
	// 但是有一种情况是例外。
	// 那就是当 case 使用关键字 fallthrough 开启穿透能力的时候。
	// fallthrough 只能穿透一层，意思是它让你直接执行下一个case的语句，而且不需要判断条件。
	s := "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world111")
	case true:
		fmt.Println("2222")
	}
}

// 判断一个同学是否有挂科记录的函数
// 返回值是布尔类型
func getResult(args ...int) bool {
	for _, i := range args {
		if i < 60 {
			return false
		}
	}
	return true
}

// for 循环
func fun3() {

	// for [condition |  ( init; condition; increment ) | Range | 啥也不接]
	// {
	// statement(s);
	// }

	//接一个条件表达式
	a := 1
	for a <= 5 {
		fmt.Println(a)
		a++
	}
	// 接三个表达式
	// 第一个表达式：初始化控制变量，在整个循环生命周期内，只运行一次；
	// 第二个表达式：设置循环控制条件，当返回true，继续循环，返回false，结束循环；
	// 第三个表达式：每次循完开始（除第一次）时，给控制变量增量或减量。
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 接 for-range 语句
	// range 会返回两个值：索引和数据，若后面的代码用不到索引，需要使用 _ 表示
	myarr := [...]string{"world", "python", "go"}
	for _, item := range myarr {
		fmt.Printf("helloxx===, %s\n", item)
	}

	//不接表达式：无限循环
	// for {
	// 	代码块
	// }

	// // 等价于
	// for ;; {
	// 	代码块
	// }
	var i int = 1
	for {
		if i > 5 {
			break
		}
		fmt.Printf("hello, %d\n", i)
		i++
	}
}

// goto 无条件跳转  java中也有 因为会破坏结构的可读性，很不常用
func fun4() {
	//goto 可以打破原有代码执行顺序，直接跳转到某一行执行代码。
	// 	goto flag
	// 	fmt.Println("B")
	// flag:
	// 	fmt.Println("A")

	// goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
	// 	i := 1
	// flag:
	// 	if i <= 5 {
	// 		fmt.Println(i)
	// 		i++
	// 		goto flag
	// 	}

	//goto语句与标签之间不能有变量声明，否则编译错误。
	// 	fmt.Println("start")
	// 	goto flag
	// 	var say = "hello oldboy"
	// 	fmt.Println(say)
	// flag:
	// 	fmt.Println("end")
}

// defer 延迟执行
// defer 的用法很简单，只要在后面跟一个函数的调用，就能实现将这个 xxx 函数的调用延迟到当前函数执行完后再执行。
// 多个defer时，类似栈一样，后进先出。
func fun5() {
	// defer myfunc()
	// fmt.Println("A")

	//换个写法
	defer fmt.Println("B")
	fmt.Println("A")

	//使用 defer 只是延时调用函数，此时传递给函数里的变量，不应该受到后续程序的影响。
	name := "go"
	defer fmt.Println(name) // 输出: go
	name = "python"
	fmt.Println(name) // 输出: python

	// 看起来可能会失效的情况
	defer2()

	// defer 是return 后才调用的  感觉相当于java里的finally执行
	myname := defer1()
	fmt.Println("main 函数里的name: ", name)
	fmt.Println("main 函数里的myname: ", myname)

	// 	并非defer在return后调用
	// return不是原子性操作，1先复制 2在调用ret指令，而defer就在 这 1 2之间，不相信的话 使用下具名函数试试：
	test()
}

func test() (x int) {
	x = 10
	defer func() {
		x++
	}()
	return x
}
func myfunc() {
	fmt.Println("B")
}

var name string = "go"

func defer1() string {
	defer func() {
		name = "python"
	}()
	fmt.Printf("defer1 函数里的name：%s\n", name)
	return name
}

func defer2() {

	//defer 后接的是一个函数，并且不是普通函数，而是一个匿名的闭包函数
	//闭包的特性，实际上在闭包函数存的是 age 这个变量的指针（

	// 	若 defer 后接的是单行表达式，那defer 中的 age 只是拷贝了 func1 函数栈中 defer 之前的 age 的值；
	// 若 defer 后接的是闭包函数，那defer 中的 age 只是存储的是 func1 函数栈中 age 的指针。
	age := 0
	defer func() {
		fmt.Println("age的值", age)
	}()
	age = 18
	return
}

// select-case
// 跟 switch-case 相比，select-case 用法比较单一，它仅能用于 信道/通道 的相关操作。
// select 随机性 不是按顺序执行

// select 只能用于 channel 的操作(写入/读出/关闭)，而 switch 则更通用一些；
// select 的 case 是随机的，而 switch 里的 case 是顺序执行；
// select 要注意避免出现死锁，同时也可以自行实现超时机制；
// select 里没有类似 switch 里的 fallthrough 的用法；
// select 不能像 switch 一样接函数或其他表达式。

func fun6() {
	//		 select {
	//	    case 表达式1:
	//	        <code>
	//	    case 表达式2:
	//	        <code>
	//	  default:
	//	    <code>
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	c1 <- "hello"

	select {
	case msg1 := <-c1:
		fmt.Println("c1 received: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 received: ", msg2)
	default:
		fmt.Println("No data received.")
	}

	// 避免造成死锁
	//select 在执行过程中，必须命中其中的某一分支,如果你没有写 default 分支，select 就会阻塞
	// 一个是，养成好习惯，在 select 的时候，也写好 default 分支代码
	// 另一个是，让其中某一个信道可以接收到数据

	// 当 case 里的信道始终没有接收到数据时，而且也没有 default 语句时，select 整体就会阻塞，
	// 但是有时我们并不希望 select 一直阻塞下去，这时候就可以手动设置一个超时时间。
	c3 := make(chan string, 1)
	c4 := make(chan string, 1)
	timeout := make(chan bool, 1)

	go makeTimeout(timeout, 2)

	select {
	case msg3 := <-c3:
		fmt.Println("c3 received: ", msg3)
	case msg4 := <-c4:
		fmt.Println("c4 received: ", msg4)
	case <-timeout:
		fmt.Println("Timeout, exit.")
	}
}
func makeTimeout(ch chan bool, t int) {
	time.Sleep(time.Second * time.Duration(t))
	ch <- true
}

// 异常机制：panic 和 recover
// panic：抛出异常，使程序崩溃
// recover：捕获异常，恢复程序或做收尾工作
// revocer 调用后，抛出的 panic 将会在此处终结，不会再外抛，但是 recover，并不能任意使用，它有强制要求，必须得在 defer 下才能发挥用途。
func fun7() {
	// 手动触发宕机
	// panic("crash")

	set_data(20)
	// 如果能执行到这句，说明panic被捕获了
	// 后续的程序能继续运行
	fmt.Println("everything is ok")
	panicTest()
}

func set_data(x int) {
	defer func() {
		// recover() 可以将捕获到的panic信息打印
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 故意制造数组越界，触发 panic
	var arr [10]int
	arr[x] = 88
}

func panicTest() {
	// 这个 defer 并不会执行
	defer fmt.Println("in main")

	go func() {
		defer println("in goroutine")
		panic("xx")
	}()

	time.Sleep(2 * time.Second)
}
