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

	cratesPiles := make([][]rune, 9)

	lineCount := 0
	readingTable := true
	moveCount := 1

	for scanner.Scan() {
		line := scanner.Text()

		if readingTable {
			for i, char := range line {
				if char == '[' {
					pos := int(math.Floor(float64(i / 4)))                       // New [letter] every 4 chars, so i / 4 == cratePiles[n]
					cratesPiles[pos] = append(cratesPiles[pos], rune(line[i+1])) // i+1 is the letter
				}
				if char == '1' && i == 1 {
					readingTable = false
					reverse(cratesPiles)
					break
				}
			}
			lineCount++
		}

		if !readingTable {

			if len(line) == 0 || line[0] != 'm' {
				continue
			}

			move, _ := strconv.Atoi(numAfterWord(line, "move"))
			from, _ := strconv.Atoi(numAfterWord(line, "from"))
			to, _ := strconv.Atoi(numAfterWord(line, "to"))

			pileFrom := cratesPiles[from-1]

			elementsFrom := []rune{}
			for i := 1; i <= move; i++ {
				elementsFrom = append(elementsFrom, pileFrom[len(pileFrom)-i])
			}

			cratesPiles[to-1] = append(cratesPiles[to-1], elementsFrom...) // Appends n element of "from" to "to" (n being "move")
			cratesPiles[from-1] = pileFrom[:(len(pileFrom) - move)]        // Reassign the origin pile with a copy of itself - cutted elements

			//fmt.Println("cratesPiles[", to-1, "] == ", string(cratesPiles[to-1]))

			moveCount++
		}
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err().Error())
	}

	printCrates(cratesPiles)
	fmt.Print("\n\n")
	printTopChars(cratesPiles)
}

