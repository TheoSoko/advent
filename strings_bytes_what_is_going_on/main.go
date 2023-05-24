package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func escapeThis(str string) string {
	replacers := map[byte]string{ //Base8
		012: "\\n",
		015: "\\r",
	}
	for i := 0; i < len(str); i++ {
		byte := str[i]
		replacer, match := replacers[byte]
		if match {
			sloice := strings.Split(str, "")
			first := append(sloice[:i], replacer)
			sloice = append(first, sloice[i+1:]...)
			str = strings.Join(sloice, "")
		}
	}
	return str
}

func main() {
	fmt.Println("Hi")
	fmt.Println("Enter some text :")
	reader := bufio.NewReader(os.Stdin)

	for true {
		str, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if str == "exit\r\n" {
			fmt.Println("Goodbye !")
			return
		}

		fmt.Println("Length of string as bytes :", len(str))
		fmt.Println("Length of string as runes (32int) :", len([]rune(str)))
		fmt.Println("Length of space trimed bytes :", len([]rune(strings.TrimSpace(str))))
		fmt.Println("Length of bytes without CR and LF :", len(strings.TrimSuffix(str, "\r\n")))
		fmt.Println()
		fmt.Println("Your string as utf-8 interpreted:", str)
		fmt.Println("Your string as runes (32int):", []rune(str))
		lastChar, beforeLast := "\\n", "\\r"
		fmt.Println("The last two characters (escaped):", lastChar, beforeLast)

		// ♪ ♫ ♪
		fmt.Println("Your string, really escaped: ", escapeThis(str))
		fmt.Println("Type \"exit\" if you wish to exit.")
		//fmt.Println(escapeThis("Hi, blablabla, I am a \n plumber, and I \r don't care !"))
	}

}
