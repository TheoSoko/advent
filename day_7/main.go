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
	SmallDirs  map[string]int
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
	nav := Navigating{
		DirList:    []Dir{{Name: "root", Size: 0, Contains: make(map[string]Dir)}},
		CurrentDir: Dir{Name: "root", Size: 0, id: uuid.New(), Contains: make(map[string]Dir)},
	}

	debugCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		nav.getnavItem(line)

		debugCount++
	}
/*
	for _, d := range nav.DirList {
		fmt.Println("Dir ", d.Name)
		fmt.Println("Size ", d.Size)
		fmt.Println()
	}
*/
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
		if itemType == "cd" {
			nav.updateDirList(nav.CurrentDir) // UPDATE DIRECTORY LIST

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
			dest, ok := nav.getDir(dirName, nav.CurrentDir)
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

func (nav *Navigating) findParentDir(current Dir) (Dir, error) {
	for _, dir := range nav.DirList {
		_, ok := dir.Contains[current.Name]
		if ok && current.Parent == dir.id {
			return dir, nil
		}
	}

	return Dir{}, fmt.Errorf(fmt.Sprint("Le dossier ", current.Name, " n'a pas de parent"))
}

func (nav *Navigating) dirExists(d Dir) bool {
	for _, dir := range nav.DirList {
		if dir.Name == d.Name && dir.Size == d.Size && dir.Parent == d.Parent {
			return true
		}
	}
	return false
}

func (nav *Navigating) findChildDirs(dirName string) map[string]Dir {
	for _, dir := range nav.DirList {
		if dir.Name == dirName {
			return dir.Contains
		}
	}

	return make(map[string]Dir)
}

func (nav *Navigating) getDir(dirName string, current Dir) (Dir, bool) {
	for _, dir := range nav.DirList {
		if dir.Name == dirName && dir.Parent == current.id {
			return dir, true
		}
	}

	return Dir{}, false
}

func (nav *Navigating) updateDirList(d Dir) {
	for _, dir := range nav.DirList {
		if dir.id == d.id {
			dir = d
		}
	}

	nav.DirList = append(nav.DirList, d)
}

/*childDirs := nav.findChildDirs("lrrl")
for _, childD := range childDirs {
	fmt.Println("dir \"lrrl\" contains ", childD.Name) // lrrl should contain dcfmtw
}*/
