package main

import (
	"os"
	"path/filepath"
	"strings"
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
