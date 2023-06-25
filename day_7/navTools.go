package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

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

func getPathToFile() string {
	ex, _ := os.Executable()
	root := filepath.Dir(ex)

	filePath := root + "/file_system"
	if strings.Contains(root, "\\AppData\\Local\\Temp") {
		filePath = "./file_system"
	}

	return filePath
}

func (nav *Navigating) findChildDirs(dirId uuid.UUID) map[string]Dir {
	for _, dir := range nav.DirList {
		if dir.id == dirId {
			return dir.Contains
		}
	}

	return make(map[string]Dir)
}

func (nav *Navigating) getChildDir(dirName string, current Dir) (Dir, bool) {
	for _, dir := range nav.DirList {
		if dir.Name == dirName && dir.Parent == current.id {
			return dir, true
		}
	}

	return Dir{}, false
}

func (nav *Navigating) getRoot() Dir {
	for _, dir := range nav.DirList {
		if dir.Name == "root" {
			return dir
		}
	}

	return Dir{}
}

func (nav *Navigating) getTotalSize(dir Dir, sum int) int {
	children := nav.findChildDirs(dir.id)

	for _, subDir := range children {
		if subDir.Size <= 100000 {
			sum += subDir.Size
		}
		children2 := nav.findChildDirs(subDir.id)
		if len(children2) > 0 {
			sum += nav.getTotalSize(subDir, sum)
		}
	}

	return sum
}

func (nav *Navigating) getDirsWithName(dirName string) []Dir {
	dirs := []Dir{}
	for _, d := range nav.DirList {
		if d.Name == dirName {
			dirs = append(dirs, d)
		}
	}

	return dirs
}

func (nav *Navigating) updateDirList(d Dir) {
	for _, dir := range nav.DirList {
		if dir.id == d.id {
			dir = d
		}
	}

	nav.DirList = append(nav.DirList, d)
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
