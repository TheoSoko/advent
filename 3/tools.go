package main


import (
	"fmt"
	"unicode"
)
//The data input is made of constants
//So I set a map as a tool to check for doublons later
func getAlpha() map[rune]bool {
	alpha := make(map[rune]bool, 52)
	for i := 0; i < 26; i++ {
		alpha[rune(i+65)] = false
		alpha[rune(i+97)] = false
	}
	return alpha
}

func getSumValues(oddOnes []rune) int{
	// Getting the priority sum thing
	sum := 0
	for _, rune := range oddOnes {
		if unicode.IsUpper(rune) {
			sum += int(rune - 38)
			fmt.Println("char: ", string(rune), "priority (upper)", int(rune-38))
			continue
		}
		sum += int(rune - 96)
		fmt.Println("char: ", string(rune), "priority (lower)", int(rune-96))
	}
	return sum
}


