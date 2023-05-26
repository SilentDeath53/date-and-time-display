package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Date and Time:")
	fmt.Println(currentTime.Format("January 02, 2006 - 15:04:05"))
}
