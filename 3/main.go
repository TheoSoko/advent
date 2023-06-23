package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("C://go-projects/advent/3/ruckstacks.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	//List of the items mistakenly placed in both compartments
	var oddOnes []rune

	//Scanning each line of text and finding the doublon
	scanner := bufio.NewScanner(file)
	//i := 0
	for scanner.Scan() {
		line := scanner.Text()
		half := len(line) / 2
		ruckstack1 := []rune{}
		ruckstack2 := []rune{}
		//Dividing one line input in two lists
		for pos, char := range line {
			if pos < half {
				ruckstack1 = append(ruckstack1, char)
			} else {
				ruckstack2 = append(ruckstack2, char)
			}
			/*
				if i == 0 {
					fmt.Println("char ==", char , "pos == ", pos, "under-half : ", pos < half)
				}
			*/
		}

		//i++

		// Checking the items
		alpha := getAlpha()
		for _, rune := range ruckstack1 {
			_, here := alpha[rune]
			if here {
				alpha[rune] = true
			}
		}
		for _, rune := range ruckstack2 {
			if alpha[rune] == true {
				oddOnes = append(oddOnes, rune)
				break // Once you find a doublon, go over the next entry
			}
		}
	}

	fmt.Println("somme des priorités: ", getSumValues(oddOnes))

}
