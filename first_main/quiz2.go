package main

import (
	"fmt"
)

func quiz2() {

	for i := 0; i <= 100; i++ {
		fmt.Printf("%d:\n", i)
		if i%3 == 0 {
			fmt.Print("Fizz")
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
		}
		fmt.Println()
	}
}
