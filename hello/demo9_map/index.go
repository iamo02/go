package main

import "fmt"

func main() {
	var number = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Print("", number["one"])

	var colre = make(map[string]string)
	colre["red"] = "#F00"
	colre["green"] = "#0F0"
	colre["blue"] = "#00F"

	fmt.Println("", colre["red"])

	var courses = make(map[string]map[string]int)

	courses["android"] = make(map[string]int)
	courses["android"]["price"] = 200
	courses["android"]["code"] = 1234

	courses["ios"] = make(map[string]int)
	courses["ios"]["price"] = 100
	courses["ios"]["code"] = 4576
	fmt.Println(courses["android"]["price"])
}
