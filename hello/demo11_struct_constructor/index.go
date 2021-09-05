package main

import (
	"demo11/employee"
)

func main() {
	e := employee.New("Pongchai", "Boonmee", 30, 20)
	e.LeavesRemaining()
}
