package main

import (
	// 导入包 会按照深度优先的顺序去执行导入包的初始化
	"github.com/astaxie/beego"
)

// 定义一个结构 struct
// 利用了 Go 语言的组合的概念，匿名包含了 beego.Controller
// 这样我们的 MainController 就拥有了 beego.Controller 的所有方法
type MainController struct {
	beego.Controller
}

// 定义了 MainController 的 Get 方法用来重写继承的 Get 函数
func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

// main 函数作为入口
func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
