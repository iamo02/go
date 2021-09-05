package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)

	ch <- 1
	fmt.Println("ch 1")

	fmt.Println(<-ch)

	ch <- 2
	fmt.Println("ch 2")

	fmt.Println(<-ch)

	time.Sleep(1 * time.Second)
}
