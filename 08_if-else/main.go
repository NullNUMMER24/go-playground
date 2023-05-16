package main

import (
	"fmt"
)

func main() {
	x := 7
	y := "test"

	if x > 8 {
		fmt.Println(y + " passed")
	} else {
		fmt.Println(y + " failed")
	}
}
