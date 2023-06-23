package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	var err error

	log.Println("Hello")
	crates, err := os.Open("C://go-projects/advent/day_5/crates.txt")
	if err != nil {
		panic(err)
	}
	defer crates.Close()

	scanner := bufio.NewScanner(crates)

	//cratesPiles := make([][8]rune, 9)
	cratesPiles := makeSlice()
	//cratesPilesR := makeSlice()
	
	lineCount := 0
	readingTable := true

	for scanner.Scan() {
		line := scanner.Text()

		if readingTable {
			for i, char := range line {
				if char == '1' && i == 1 {
					readingTable = false
					break
				}
				if char == '[' {
					pos := int(math.Floor(float64(i / 4)))        // New [letter] every 4 chars, so i / 4 == cratePiles[n]
					cratesPiles[pos][lineCount] = rune(line[i+1]) // i+1 is the letter
				}
			}
			lineCount++
		}

		if !readingTable {
			reverse(cratesPiles)

			if len(line) == 0 || line[0] != 'm' {
				continue
			}

			move, _ := strconv.Atoi(string(line[indexAfterWord(line, "move")]))
			intFrom, _ := strconv.Atoi(string(line[indexAfterWord(line, "from")]))
			intTo, _ := strconv.Atoi(string(line[indexAfterWord(line, "to")]))

			fmt.Sprintln("move: ", move)
			fmt.Sprintln("from: ", intFrom)
			fmt.Sprintln("to: ", intTo)

			to := cratesPiles[intTo - 1]
			from := cratesPiles[intFrom - 1]

			cratesPiles[intTo - 1] = append(to, from[len(from)-1])
			
			//fmt.Println("Moved last of cratesPilesR[", )
		}
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err().Error())
	}

		/*
		fmt.Print("cratesPiles[0][0] : ", string(cratesPiles[0][0]), "\n")
		fmt.Print("cratesPiles[0][7] : ", string(cratesPiles[0][7]), "\n")
		fmt.Print("cratesPiles[8][0] : ", string(cratesPiles[8][0]), "\n")
		fmt.Print("cratesPiles[8][7] : ", string(cratesPiles[8][7]), "\n")
		fmt.Println()
		*/
		/*
		fmt.Print("cratesPilesR[0][0] : ", string(cratesPilesR[0][0]), "\n")
		fmt.Print("cratesPilesR[0][7] : ", string(cratesPilesR[0][7]), "\n")
		fmt.Print("cratesPilesR[8][0] : ", string(cratesPilesR[8][0]), "\n")
		fmt.Print("cratesPilesR[8][7] : ", string(cratesPilesR[8][7]), "\n")
		*/
}

func indexAfterWord(s string, sub string) int {
	subCounter := 0
	for i, c := range s {
		if c == rune(sub[subCounter]) {
			subCounter++
		} else {
			subCounter = 0
		}
		if subCounter == len(sub) {
			return i + 2
		}
	}
	return 0
}

func reverse(a [][]rune) [][]rune {
	for i, a2 := range a {
		for                             //
		left, right := 0, len(a2)-1;    //
		left < right;                   //
		left, right = left+1, right-1 { //
			a[i][left], a[i][right] = a2[right], a2[left]
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
