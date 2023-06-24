package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Navigating struct {
	dirList map[string]int
	isDir bool
}

func main() {
	var err error

	ex, _ := os.Executable()
	root := filepath.Dir(ex)
	filePath := root + "/file_system"
	if strings.Contains(root, "\\AppData\\Local\\Temp") {
		filePath = "./file_system"
	}

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		line := scanner.Text()

		navigate(line)
	}

}


func navigate(l string){
	nav := Navigating{}
	nav.isDir = false

	itemType, pos := extractStr(l, 0)
	//fmt.Println("itemType : ", itemType)
	if itemType == "Dir" {
		nav.isDir = true
		dirName, _ := extractStr(l, pos+len("dir"))
		nav.dirList[dirName] = 45
	}

	fmt.Println("directories[`\"gts\"]", nav.dirList["gts"])
}

func extractStr(s string, pos int) (sub string, posOut int) {
	sub = ""
	on := false
	for i := pos; i < len(s); i++ {
		if on && s[i] == ' ' {
			return sub, i
		}
		if s[i] != ' ' {
			on = true
			sub += string(s[i])
		}
	}
	return sub, -1
}
