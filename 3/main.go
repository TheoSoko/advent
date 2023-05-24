package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func randKeyInMap[Comp comparable, a any](m map[Comp]a) Comp {
	randomInt := rand.Intn(len(m))
	iter := 0
	for key := range m {
		if iter == randomInt {
			return key
		}
		iter++
	}
	panic("aaaaaah")
}

func main() {

	type Choice struct {
		beatenBy string
		winsOver string
		userChar string
		compChar string
	}

	choices := map[string]Choice{
		"pierre": {
			beatenBy: "feuille",
			winsOver: "ciseaux",
			userChar: "X",
			compChar: "A",
		},
		"feuille": {
			beatenBy: "ciseaux",
			winsOver: "pierre",
			userChar: "Y",
			compChar: "B",
		},
		"ciseaux": {
			beatenBy: "pierre",
			winsOver: "feuille",
			userChar: "Z",
			compChar: "C",
		},
	}

	points := map[string]int{
		"lose":    0,
		"draw":    3,
		"win":     6,
		"pierre":  1,
		"feuille": 2,
		"ciseaux": 3,
	}

	gameEval := func(uChoice string, compChoice string) (winner string, userPoints int) {
		if uChoice == compChoice {
			return "égalité", points["draw"] + points[uChoice]
		}
		if choices[uChoice].beatenBy == compChoice {
			return "machine", points["lose"] + points[uChoice]
		}
		if choices[uChoice].winsOver == compChoice {
			return "utilisateur", points["win"] + points[uChoice]
		}
		panic("Problème inconnu à gameEval")
	}

	//score1 := 0
	//score2 := 0.....

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Entrez votre choix (pierre, feuille, ciseaux)")
	var userChoice string

	for true {
		read, err := reader.ReadString('\n')
		userChoice = strings.TrimSpace(read)
		if err != nil {
			panic(err)
		}

		_, ok := choices[userChoice]
		if ok {
			break
		}
		fmt.Println("Veuillez choisir entre \"pierre\", \"feuille\", et \"ciseaux\"")
	}

	fmt.Println("\n Aucune importance, le jeu se déroulera selon le \"guide de stratégie\" de toutes façons 🥴")

	file, err := os.Open("C:/go-projects/advent/3/strategy_guide.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Err à l'ouverture du fichier")
		panic(err)
	}

	totalPoints := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		moves := strings.Fields(line)
		var userMove string
		var compMove string
		for key, val := range choices {
			if val.compChar == moves[0] {
				compMove = key
			}
			if val.userChar == moves[1] {
				userMove = key
			}
		}
		_, userPoints := gameEval(userMove, compMove)
		totalPoints += userPoints
	}

	//computerChoice := randKeyInMap(choices)

	//winner, userPoints := gameEval(userChoice, computerChoice)
	//fmt.Println("Le gagnant est : ", winner, "!")
	fmt.Println("Vous avez :", totalPoints, "points au total !")
}
