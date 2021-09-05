package main

import (
	"fmt"
	"time"
)

func main() {

	go run1()
	go run2()

	time.Sleep(5 * time.Second)
}

func run1() {
	for i := 0; i < 50; i++ {
		fmt.Println("run  1", i)
	}

}

func run2() {

	for i := 0; i < 50; i++ {
		fmt.Println("run  2", i)
	}
}
