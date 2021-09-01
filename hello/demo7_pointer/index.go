package main

import "fmt"

func main() {
	msg := "some message"
	var msgPointer *string = &msg
	fmt.Println(msg)
	fmt.Println(*msgPointer)

	chengMessage(&msg, "new Message1")
	fmt.Println(msg)

	chengMessage(msgPointer, "new Message")
	fmt.Println(msg)
	fmt.Println(*msgPointer)
}

func chengMessage(aPointer *string, newMessage string) {
	*aPointer = newMessage
}
