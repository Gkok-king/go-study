package main

import "fmt"

func main() {
	// 实例化
	// myself := Profile{name: "小明", age: 24, gender: "male"}
	// // 调用函数
	// myself.FmtProfile()
	// myself.increase_age()
	// myself.FmtProfile()
	fun4()
}

// 1定义一个名为Profile 的结构体
type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile // 指针
	father *Profile // 指针

	// 若相邻的属性（字段）是相同类型，可以合并写在一起
	phone, address string

	// 规则一：当最后一个字段和结果不在同一行时，, 不可省略。反之，不在同一行，就可以省略。
	// 规则二：字段名要嘛全写，要嘛全不写，不能有的写，有的不写。
	// 规则三：初始化结构体，并不一定要所有字段都赋值，未被赋值的字段，会自动赋值为其类型的零值。
}

// 2定义一个与 Profile 的绑定的方法
func (person Profile) FmtProfile() {
	fmt.Printf("名字：%s\n", person.name)
	fmt.Printf("年龄：%d\n", person.age)
	fmt.Printf("性别：%s\n", person.gender)
}

// 3方法的参数传递方式
// 重点在于这个星号: *
func (person *Profile) increase_age() {
	person.age += 1
	fmt.Println("增加年龄===============")
}

// 两种定义方法的方式：
// 1以值做为方法接收者
// 2以指针做为方法接收者
// 那我们如何进行选择呢？以下几种情况，应当直接使用指针做为方法的接收者。
// 1你需要在方法内部改变结构体内容的时候
// 2出于性能的问题，当结构体过大的时候

//4实现继承 Go 语言本身并不支持继承。可以使用组合的方法，实现类似继承的效果。

type staff struct {
	name     string
	age      int
	gender   string
	position string
	company  // 匿名字段
}

type company struct {
	companyName string
	companyAddr string
}

// 5内部方法与外部方法
// 在 Go 语言中，函数名的首字母大小写非常重要，它被来实现控制对方法的访问权限。
// 当方法的首字母为大写时，这个方法对于所有包都是Public，其他包可以随意调用
// 当方法的首字母为小写时，这个方法是Private，其他包是无法访问的。
func Join() {

}

func join() {

}

// 6. 三种实例化方法
func fun1() {
	xm := Profile{
		name:   "小明",
		age:    18,
		gender: "male",
	}
	fmt.Println(xm)
}

// 使用 new
func fun2() {
	xm := new(Profile)
	// 等价于: var xm *Profile = new(Profile)
	fmt.Println(xm)
	// output: &{ 0 }

	xm.name = "iswbm"  // 或者 (*xm).name = "iswbm"
	xm.age = 18        //  或者 (*xm).age = 18
	xm.gender = "male" // 或者 (*xm).gender = "male"
	fmt.Println(xm)
	//output: &{iswbm 18 male}
}

// 使用 &
func fun3() {
	var xm *Profile = &Profile{}
	fmt.Println(xm)
	// output: &{ 0 }

	xm.name = "iswbm"  // 或者 (*xm).name = "iswbm"
	xm.age = 18        //  或者 (*xm).age = 18
	xm.gender = "male" // 或者 (*xm).gender = "male"
	fmt.Println(xm)
	//output: &{iswbm 18 male}
}

// 7选择器
func fun4() {
	//当对象是结构体对象的指针时，你想要获取字段属性时
	p1 := &Profile{name: "小明", age: 24, gender: "male"}
	fmt.Println((*p1).name) // output: iswbm

	// 但还有一个更简洁的做法，可以直接省去 * 取值的操作，选择器 . 会直接解引用，示例如下
	fmt.Println(p1.name)
}
