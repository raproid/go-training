package main

import (
	    "fmt"
	    "log"
)

func panicker_example_with_handling_the_panic() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")

	func panicker() {
		fmt.Println("about to panic")
		defer func() {
			if err := recover(); err != nil {
				log.Println("Error:", err)
			}
		}()
		panic("panicking")
		fmt.Println("done panicking")
	}
}

