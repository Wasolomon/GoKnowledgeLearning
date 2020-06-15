package main

import "fmt"

func stackingDefers()  {
	defer func() {
		fmt.Printf("1")
	}()
	defer func() {
		fmt.Printf("2")
	}()
	defer func() {
		fmt.Printf("3")
	}()
}

// 匿名返回值
func returnValues() int {
	var result int
	defer func() {
		// 返回刚才创建的返回值（retValue）
		result++
		fmt.Println("defer")
	}()
	// 将result赋值给返回值 可以理解成Go自动创建了一个返回值retValue，相当于执行retValue = result
	// 检查是否有defer，如果有则执行
	return result
}

// 命名返回值
func namedReturnValues() (result int) {
	// return和defer是“同时”执行的
	defer func() {
		result++
		fmt.Println("defer")
	}()
	// 由于返回值在方法定义时已经被定义，所以没有创建retValue的过程
	// result就是retValue，defer对于result的修改也会被直接返回
	return result
}

func main()  {
	stackingDefers()	// 执行结果为 321

	r := returnValues()
	fmt.Println(r)	// 0

	k := namedReturnValues()
	fmt.Println(k)	// 1
}
