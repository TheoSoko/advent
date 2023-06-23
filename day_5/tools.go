package main

import (
	"fmt"
	"strconv"
)

func numAfterWord(s string, sub string) string {
	var numIndex int
	numStr := ""
	subCounter := 0

	for i, c := range s {
		if c == rune(sub[subCounter]) {
			subCounter++
		} else {
			subCounter = 0
		}
		if subCounter == len(sub) {
			numIndex = i + 2
			break
		}
		if i == len(s)-1 {
			return ""
		}
	}

	for i := numIndex; i < len(s); i++ {
		charByte := s[i]
		_, err := strconv.Atoi(string(charByte))
		if err != nil {
			return numStr
		}
		numStr += string(charByte)
	}
	return numStr
}

func reverse(a [][]rune) [][]rune {
	for i, a2 := range a {
		for                             //
		left, right := 0, len(a2)-1;    //
		left < right;                   //
		left, right = left+1, right-1 { //
			a2[left], a2[right] = a2[right], a2[left]
		}
		a[i] = a2
	}
	return a
}

func makeSlice() [][]rune {
	sliceWithAlloc := make([]rune, 8, 100)
	slice := make([][]rune, 9)
	for i := range slice {
		slice[i] = sliceWithAlloc
	}
	return slice
}

func printTopChars(a [][]rune) {
	topChars := ""
	for _, a2 := range a {
		if len(a2) > 0 {
			topChars += string(a2[len(a2)-1])
		}
	}
	fmt.Println(topChars)
}

func printCrates(crates [][]rune) {
	longestPile := 0
	for i := 1; i <= len(crates); i++ {
		fmt.Print(" ", i, "  ")
		if len(crates[i-1]) > longestPile {
			longestPile = len(crates[i-1])
		}
	}
	fmt.Println()
	for i := 0; i < longestPile; i++ {
		for _, pile := range crates {
			if i >= len(pile) {
				fmt.Print("    ")
				continue
			}
			fmt.Print("[", string(pile[i]), "] ")
		}
		fmt.Println()
	}
}
