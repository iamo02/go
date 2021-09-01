package main

import "fmt"

func main() {
	fnFor()
	fnWile()
	for_range()
	foruserbreak()
}

func fnFor() {
	for i := 0; i <= 10; i++ {
		fmt.Println("for : ", i)
	}
}

func fnWile() {

	var index int = 0
	for index < 5 {

		fmt.Println("while", index)
		index++
	}
}

func for_range() {
	index := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for i, item := range index {
		fmt.Printf("%d %d \n", i, item)
	}

	for _, item := range index {
		fmt.Println(item)
	}
}

func foruserbreak() {
	i := 0
	for true {
		i++
		if i > 5 {
			break
		}
		fmt.Println("break : ", i)
	}
}
