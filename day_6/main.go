package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var err error

	ex, _ := os.Executable()
	root := filepath.Dir(ex)

	filePath := root+"/datastream_buffer"
	if strings.Contains(root, "\\AppData\\Local\\Temp") {
		filePath = "./datastream_buffer"
	}

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		line := scanner.Text()
		marker := findMarker(line, 14)

		if marker == 0 {
			fmt.Println("Echec, marqueur non initialisé")
			return
		}

		//subStr := line[marker-14:]

		fmt.Println("marker == ", marker)
		//fmt.Println("subStr == ", subStr)
		return
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err.Error())
	}

}

func findMarker(s string, markerSignal int) int {
	charBuff := ""
	var marker int

	for i, c := range s {
		charBuff += string(c)

		if findDuplicate(charBuff) {
			charBuff = string(c)
		}

		if len(charBuff) == markerSignal {
			fmt.Println("hey")
			marker = i + 1 // range starts at 0
			break
		}
	}
	fmt.Println(charBuff)

	return marker
}

func findDuplicate(s string) bool {
	for i, c := range s {
		for i2, c2 := range s {
			if c == c2 && i != i2 {
				return true
			}
		}
	}
	return false
}
