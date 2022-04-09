package main

import (
	"bufio"
	"fmt"
	"hangman/Utils"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {

	if len(os.Args) != 2 {
		return
	}

	rand.Seed(time.Now().UnixNano())

	data := Utils.HangManData{Attempts: 10, Error: ""}
	data.Word = Utils.GetRandomWord(os.Args[1])
	data.ToFind = strings.Repeat("_", utf8.RuneCountInString(data.Word))
	data.HangmanPositions = Utils.ParseHangmanFile("./hangman.txt")

	Utils.RevealRandomLetter(&data)

	fmt.Printf("Good luck you have %d attempts !\n%s\n", data.Attempts, data.ToFind)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nChoose: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			s := Utils.HangMan(&data, text)
			fmt.Print(s)
			fmt.Println(data.ToFind)
			if 10-data.Attempts-1 >= 0 {
				fmt.Print(data.HangmanPositions[10-data.Attempts-1])
			}
			if data.Word == data.ToFind || data.Attempts <= 0 {
				if data.Word == data.ToFind {
					fmt.Print("\nCongrats !")
				} else {
					fmt.Printf("\nYou loose.., the word was %s", data.Word)
				}
				fmt.Println()
				break
			}
		}

	}
}
