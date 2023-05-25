package main

import (
	"fmt"
	//"math"
)

func main() {
	size := 80
	start := size / 2
	end := size / 2

	for true {
		for i := 0; i <= size; i++ {
			if i >= start && i <= end {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		start -= 2
		end += 2
		if start == 0 {
			return
		}
		fmt.Println()
	}

}
