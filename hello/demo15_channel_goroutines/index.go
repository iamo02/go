package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	go routine1(ch, 1)

	fmt.Println(<-ch)
	time.Sleep(1 * time.Second)
}

func routine1(c chan int, payloadd int) {
	c <- payloadd
	fmt.Println("print 1")
	c <- payloadd
	fmt.Println("print 2")
	c <- payloadd
	fmt.Println("print 3")
}
