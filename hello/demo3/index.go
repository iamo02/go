package main

import (
	"fmt"
)

var count int = 0

func main() {
	fmt.Println("")

	var tmp1 int = 0
	tmp1 = 10
	var tmp2 string = "hello"
	var tmp3 bool = true

	const temp int = 10
	// temp = 2

	fmt.Print(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)
	fmt.Println(temp)

	temp5 := 20 // var temp int = 20
	fmt.Print(temp5)

	run()
}

func run() {
	count++
	fmt.Println(count)
}
