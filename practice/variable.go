package main

import "fmt"

func main(){
	fmt.Println("ddd")
	// var 변수명 자료형   <- 명시적
	var test1 string = "string"
	var test2 string
	// 변수명 :=    <-타입 추론해줌
	test3 := "33"
	test4 := 44

	// var 변수명 =     이렇게 하면 묵시적

	test2 = "string2"

	//밑에 세개가 다 된다.
	var t1 string = "dd"
	var t2  = "dd"
	t3 := "dd"
	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
	fmt.Println(test4)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)

}
