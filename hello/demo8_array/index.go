package main

import "fmt"

func main() {
	var array1 []int = []int{1, 3, 4}
	var array2 = []int{0, 9, 8, 7}
	var array3 [3]string

	for _, item := range array1 {
		fmt.Println(item)
	}

	for _, item := range array2 {
		fmt.Println(item)
	}

	array3[0], array3[1], array3[2] = "Pongchai", "Boonmee", "O"
	for _, item := range array3 {
		fmt.Println(item)
	}

	//var number = make([]int, 3, 5)
	var number = []int{1, 2, 3}
	number = append(number, 4)
	number = append(number, 5)
	number = append(number, 6)
	showSlice(number)

	number = number[1:len(number)] //2,3,4,5,6
	showSlice(number)

	number = number[0 : len(number)-1] //2,3,4,5
	showSlice(number)

	number = removeIndex(number, 2)
	showSlice(number)

}

func showSlice(x []int) {
	fmt.Printf("len= %d cap= %d slice= %v \n", len(x), cap(x), x)
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)

}
