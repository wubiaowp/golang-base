package main

import "fmt"

/**
 Go语言接口使用案例
 */
 type Phone interface {
 	call()
 }

 type Nokia struct {

 }

 func (nokia Nokia) call(){
 	fmt.Print("this is nokia\n")
 }

 type Iphone struct {

 }

 func(iphone Iphone) call(){
 	fmt.Print("this is iphone")
 }

func main()  {
	var phone Phone

	phone = new(Nokia)
	phone.call()

	phone = new(Iphone)
	phone.call()

}