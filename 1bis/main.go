package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	//"os"
	//"bufio"
)

func main() {
	//cd := exec.Command("cd", "../1")
	run := exec.Command("../1/1.exe")

	result, err := run.Output()
	if err != nil {
		panic(err)
	}

	sloice := strings.Fields(string(result))
	threeMost := []int{0, 0, 0}
	for _, line := range sloice {
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Print("could not convert : \"", line, "\"\n")
			continue
		}
		if num > threeMost[0] {
			threeMost[0] = num
		} else if num > threeMost[1] {
			threeMost[1] = num
		} else if num > threeMost[2] {
			threeMost[2] = num
		}
	}

	total := threeMost[0] + threeMost[1] + threeMost[2]
	fmt.Println("three best : \"", threeMost, "\"")
	fmt.Println("Total: : \"", total, "\"")
}
