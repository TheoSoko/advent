package main

import (
	"fmt"
	//"math"
)

func main() {
	size := 80
	start := size / 2
	end := size / 2

	prLine := func(size int, start int, end int) {
		for i := 0; i <= size; i++ {
			if start <= i && i <= end {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
	}

	for {
		i := 0
		for i < 2 {
			prLine(size+i*4, start+i*4, end+i*4)
			//i++
		}
		if start == 0 {
			return
		}
		start -= 2
		end += 2
		fmt.Println()
	}

}
