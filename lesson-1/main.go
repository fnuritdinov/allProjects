package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(len("Hello world"))

	//fmt.Println(string("Hello World"[0]))

	//var name string

	//fmt.Scan(&name)
	//fmt.Printf("Меня зовут: %s", name)

	//var isLogic bool
	//isLogic = true
	//fmt.Println(isLogic)

	//var symbol int32 = 'c'
	//fmt.Println(string(symbol))

	//var my_var_35 string = "Hello World"
	//fmt.Println(my_var_35)

	//var name string
	//var age int

	//fmt.Println("введите ваше имя: ")
	//fmt.Scan(&name)
	//fmt.Println("введите ваш возраст:")
	//fmt.Scan(&age)
	//fmt.Println(name, age)

	//fmt.Print("hello", 24)
	//fmt.Println("hello world", 24)

	//var a int
	//fmt.Scan(&a)
	//fmt.Println(a * a)

	//var a int
	//fmt.Scan(&a)
	//fmt.Println(a % 10)

	/*var d int
	fmt.Scan(&d)
	ch := d / 30
	ch1 := (d % 30) / 60
	fmt.Println("It is", ch, "hours", ch1, "minutes")
	*/
	x := 2
	y := 3
	lim := 10
	if v := math.Pow(float64(x), float64(y)); v < float64(lim) {
		fmt.Printf("Результат: %.2f меньше лимита", v)
	}

}
