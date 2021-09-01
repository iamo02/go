package main

import "fmt"

func main() {
	fn1()
	fn2("hello go")
	fn3("good morning", 2)
	const a, b = 2, 3
	fmt.Printf("%d + %d = %d \n", 2, 3, sum(a, b))
	var x, y int = swap(a, b)
	fmt.Printf("%d,%d => %d,%d \n", x, y, a, b)

	_x, _y, title := swapv2(10, 20)
	fmt.Printf("%d,%d ,%s\n", _x, _y, title)
}

func fn1() {
	fmt.Println("fn1")
}

func fn2(msg string) {
	fmt.Println(msg)
}

func fn3(title string, version int) {
	fmt.Println(title)
	fmt.Println(version)
}

func sum(a, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}

func swapv2(a int, b int) (x, y int, title string) {
	x = b
	y = a
	title = "Reslt"
	return
}
