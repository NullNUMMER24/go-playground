/*
Print the numbers betweeen 1 to 20. 
- If the number is devisible by 3, print "fizz" instead
- If the number is devisible by 5, print "buzz" instead
- If the number is devisible by 3 and 5, print "fizz buzz" instead
- Otherwise print the number
*/

package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
