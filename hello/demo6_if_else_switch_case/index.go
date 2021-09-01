package main

import "fmt"

func main() {
	someValue := 10

	if someValue == 10 {
		fmt.Println("=10")
	} else {
		fmt.Println("!=10")
	}

	if someValue > 10 || someValue <= 2 {
		fmt.Println("someValue > 10 || someValue <= 2")
	} else {
		fmt.Println("not someValue > 10 || someValue <= 2")
	}

	if reslut := something(); reslut == "ok" {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

	fnSwitch()

}

func something() string {

	return "ok"
}

func fnSwitch() {
	index := 2
	switch index {
	case 0:
		fmt.Println("0")
		break
	case 1:
		fmt.Println("1")

	case 2:
		fmt.Println("2")

	case 3:
		fmt.Println("3")

	default:
		fmt.Println("default")

	}
}
