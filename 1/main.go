package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
)

func main() {

	file, err := os.Open("C:/go-projects/advent/1/items.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Err à l'ouverture du fichier")
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	allKal := [][]int {}
	kals := []int {}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			allKal = append(allKal, kals)
			kals = nil //empty the list
			continue
		}
		kalInt, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Err pour convertir ", line)
			continue
		}
		kals = append(kals, kalInt)
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("Err au scan")
	}

	//fmt.Println("All the kal values in index 35 : ")
	//fmt.Println(allKal[35])

	//Sums
	var sums = make([]int, len(allKal)) 
	for index, kalBunch := range(allKal){
		sum := 0
		for _, intKal := range(kalBunch){
			sum += intKal
		}
		sums[index] = sum
	}

	//Highest sum
	highest := 0
	var hIndex int
	for index, sum := range(sums){
		if sum > highest {
			highest = sum
			hIndex = index
		}
	}
	fmt.Print("La somme la plus grande est: \"", highest, "\"\n")
	fmt.Print("L'index est: \"", hIndex, "\"\n")

	//prints the sums
	for _, num := range(sums){
		fmt.Println(num)
	}


}
