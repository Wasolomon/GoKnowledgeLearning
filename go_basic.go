package main

import (

	"fmt"
	. "solomon/github.com/GoKnowledgeLearning/stringutil"
)

func main()  {
	// 这是我的第一个程序将结果输出到控制台
	/*
	这是多行注释的写法
	这是多行注释的写法
	这是多行注释的写法
	*/
	fmt.Printf(Reverse("Hello World"))
	var a *int
	var b []int
	var c map[string] int
	var d chan int
	var e func(string) int
	var f error
	fmt.Printf("%v %v %v %v %v %v", a, b, c, d, e, f)
	const (
		a1 = iota
		b1
		c1
		d1 = "solomon"
		e1
		f1 = 100
		g1
		h1 = iota
		i1
	)
	fmt.Println(a1, b1, c1, d1, e1, f1, g1, h1, i1)

	const (
		a2 = 1 << iota
		b2 = 3 << iota
		c2
		d2
	)
	fmt.Println("a2 = ", a2)
	fmt.Println("b2 = ", b2)
	fmt.Println("c2 = ", c2)
	fmt.Println("d2 = ", d2)

	var a3 int=20
	var ip *int
	ip = &a3
	fmt.Printf("a3变量的地址是：%x\n", &a3)
	fmt.Printf("ip变量储存的指针地址是：%x\n", ip)
	fmt.Printf("ip变量的值是：%x\n", *ip)

	type Books struct {
		title string
		author string
		subject string
		bookID int
	}
	fmt.Println(Books{"Go 语言", "www.go.org", "Go 语言教程", 6495407})

	var number = make([]int, 3, 5)
	printSlice(number)

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum = ", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index = ", i)
		}
	}
	strs := map[string]string{"a": "apple", "c": "banana"}
	for k, v := range strs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

// 在终端中运行命令为 go run go_basic.go

// 标识符 由字母、数字和下划线组成

// 关键字 25个 预定义标识符 36个
/*
break default func interface select case defer go map struct chan else goto
package switch const fallthrough is range type continue for import return var
*/

/*
append bool byte cap close complex complex64 complex128 uint16
copy false float32 float64 imag int int8 int8 int16 uint32 int32
int64 iota len make new nil panic uint64 print println real recover
string true uint uint8 uintptr
*/

// 变量声明 var age int;

// Go语言数据类型
/*
	布尔型
		true 或者 false
		var b bool = true
	数字类型
		int float32 float64
		uint8 无符号8位整型
		uintptr
		complex64 complex128
	字符串类型
		一串固定长度的字符连接起来的字符序列
	派生类型
		指针类型 Pointer
		数组类型
		结构化类型 struct
		Channel 类型
		函数类型
		切片类型
		接口类型 interface
		Map类型
*/

// 变量声明
	// 如果没有初始化则变量默认为零值
		// 数值类型（包括complex64/128）为 0
		//布尔类型为 false
		//字符串为 ""（空字符串）
		//以下几种类型为 nil：
	// := 声明新的变量

// iota 特殊常量 可以被认为是一个被编译器修改的常量

// 算术运算符 + - * / % ++ --
// 逻辑运算符 &&(and) ||(or) !(not)
// 位运算符 & | ^ <<(乘以2的n次方) >>(除以2的n次方)
// *(指针变量) &(返回变量存储地址)

// 运算符优先级
	// 优先级5 * / % << >> & &^
	// 优先级4 + - | ^
	// 优先级3 == != < <= > >=
	// 优先级2 &&
	// 优先级1 ||


// 指针 一个指针变量指向了一个值的内存地址; 当一个指针被定义后没有分配到任何变量时值为 nil

// 数组可以存储同一类型的数据; 但在结构体中我们可以为不同项定义不同的数据类型

// cap() 可以测量切片最长可以达到多少

// range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素
