package main

import (
	"demo11/employee"
)

func main() {
	e := employee.Init("Pongchai", "Boonmee", 30, 20)
	e.LeavesRemaining()

	e = employee.Init("xxx", "Boonmee", 30, 20)
	e.LeavesRemaining()
}
