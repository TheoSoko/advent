package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	//"fmt"
)

func main() {
	var err error

	log.Println("Hello")
	pairs, err := os.Open("C://go-projects/advent/day_4/assignment_pairs.txt")
	if err != nil {
		panic(err)
	}
	defer pairs.Close()

	scanner := bufio.NewScanner(pairs)

	//contained := 0
	overLap := 0

	for scanner.Scan() {

		line := scanner.Text()
		assignments := strings.Split(line, ",")
		ass1 := strings.Split(assignments[0], "-")
		ass2 := strings.Split(assignments[1], "-")

		str1, str1bis := ass1[0], ass1[1]
		str2, str2bis := ass2[0], ass2[1]

		n1, err := strconv.Atoi(str1)
		n1bis, err2 := strconv.Atoi(str1bis)
		if err != nil || err2 != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(str2)
		n2bis, err2 := strconv.Atoi(str2bis)
		if err != nil || err2 != nil {
			panic(err)
		}

		if isRangeContained(n1, n1bis, n2, n2bis) || doRangesOverlap(n1, n1bis, n2, n2bis) {
			overLap++
		}
		//log.Printf(fmt.Sprint(n1)+"-"+fmt.Sprint(n1bis)+", "+fmt.Sprint(n2)+"-"+fmt.Sprint(n2bis))
	}

	if scanner.Err() != nil {
		log.Fatal("scanner err", scanner.Err())
	}

	//log.Print("Total contained ranges : ", contained)
	log.Print("Total overlapping ranges : ", overLap)

}

func isRangeContained(n1, n1b, n2, n2b int) bool {
	if n1 <= n2 && n2b <= n1b {
		return true // 2 is contained in 1
	}
	if n2 <= n1 && n1b <= n2b {
		return true // 1 is contained in 2
	}
	return false
}

func doRangesOverlap(n1, n1b, n2, n2b int) bool {
	if n1 <= n2 && n2 <= n1b { // 4-7, 3-5
		return true // n2 is between 1 and 1bis
	}
	if n2 <= n1 && n1 <= n2b { // 4-7, 3-5
		return true // n1 is between 2 and 2bis
	}
	/*
		If one is not fully contained in the other :
			If n1 is between n2 and n2b
				Then n2b is between n1 and n1b
				And vice-versa
	*/
	return false
}
