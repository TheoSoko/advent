package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
)

type Dir struct {
	Name     string
	Size     int
	Contains map[string]Dir
	Parent   uuid.UUID
	id       uuid.UUID
}

type Navigating struct {
	DirList    []Dir
	SmallDirs  []Dir
	CurrentDir Dir
}

func main() {
	var err error

	file, err := os.Open(getPathToFile())
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	navigate(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err.Error())
	}
}

func navigate(scanner *bufio.Scanner) {
	root := Dir{Name: "root", Size: 0, id: uuid.New(), Contains: make(map[string]Dir)}
	nav := Navigating{
		DirList:    []Dir{root},
		CurrentDir: root,
	}

	debugCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		nav.getnavItem(line)
		debugCount++
	}

	// getRoot()
	for _, d := range nav.DirList {
		//fmt.Println("Dir ", d.Name)
		//fmt.Println("Size ", d.Size)
		
		if nav.getTotalSize(d, d.Size) <= 100000 {
			nav.SmallDirs = append(nav.SmallDirs, d)
		}
	}

	total := 0
	for _, d := range nav.SmallDirs{
		fmt.Println("small dir", d.Name ,"size == ", d.Size)
		total += d.Size
	}

	fmt.Println("Total sum of the sizes of directories <= 100000 : ", total)
	fmt.Println("SmallDirs length == ", len(nav.SmallDirs))
	fmt.Println("DirList length == ", len(nav.DirList))
	fmt.Println()

	pcqjnl := nav.getDirsWithName("root")
	for _, d := range pcqjnl {
		fmt.Println("root.size == ", d.Size)
	}

}

func (nav *Navigating) getnavItem(line string) {
	itemType, pos := extractStr(line, 0)

	if itemType == "dir" {
		dirName, _ := extractStr(line, pos)
		nav.CurrentDir.Contains[dirName] = Dir{Name: dirName, Size: 0, id: uuid.New(), Contains: make(map[string]Dir)}
		fmt.Println("Current contains ", dirName)
		return
	}

	// if itemType == number
	size, err := strconv.Atoi(itemType)
	if err == nil {
		nav.CurrentDir.Size += size
		fmt.Println("Current dir size ", nav.CurrentDir.Size)
		return
	}

	if itemType == "$" {
		itemType, pos2 := extractStr(line, pos)
		// Change Directory
		if itemType == "cd" {
			// UPDATE DIRECTORY LIST
			nav.updateDirList(nav.CurrentDir) 

			dirName, _ := extractStr(line, pos2)

			fmt.Println("cd ", dirName)

			if dirName == ".." {
				parent, err := nav.findParentDir(nav.CurrentDir)
				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Println("moving up to ", parent.Name)
				//fmt.Println(parent.Name, " size is ", parent.Size)
				nav.CurrentDir = parent
				return
			}

			// Change or Create Directory
			dest, ok := nav.getChildDir(dirName, nav.CurrentDir)
			if ok {
				nav.CurrentDir = dest
			}
			id := uuid.New()
			nav.CurrentDir = Dir{Name: dirName, id: id, Size: 0, Parent: nav.CurrentDir.id, Contains: make(map[string]Dir)}
			nav.DirList = append(nav.DirList, nav.CurrentDir)

			fmt.Println("currentDir", nav.CurrentDir.Name, "has parent ", nav.CurrentDir.Parent)
			return
		}
	}
}




