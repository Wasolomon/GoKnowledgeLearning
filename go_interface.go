package main

import "fmt"

type Phone interface {
	call()
}

type ApplePhone struct {
}

func (applePhone ApplePhone) call() {
	fmt.Println("I am apple, I can call you")
}

type IPhone struct {
}

func (iphone IPhone) call() {
	fmt.Println("I am iphone, I can call you!")
}

func main()  {
	var phone Phone
	phone = new(ApplePhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
