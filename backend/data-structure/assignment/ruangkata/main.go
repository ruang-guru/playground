package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ruang-guru/playground/backend/data-structure/assignment/ruangkata/dictionary"
	"github.com/ruang-guru/playground/backend/data-structure/assignment/ruangkata/spellchecker"
)

func main() {
	for {
		fmt.Println("1) Definition")
		fmt.Println("2) Spellchecker")
		fmt.Println("0) Exit")

		fmt.Println()

		var choice int
		fmt.Print("Enter your choice: ")

		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		switch choice {
		case 0:
			fmt.Println("Bye")
			return
		case 1:
			err := definitionMenu()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			spellCheckerMenu()
		}

		fmt.Println()
	}
}

func definitionMenu() error {
	dict, err := dictionary.NewEnglishDictionary()
	if err != nil {
		return err
	}
	for {
		fmt.Println()
		fmt.Print("Enter a word (type 0 to exit): ")
		var word string
		_, err := fmt.Scanln(&word)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		if word == "0" {
			return nil
		}

		definition, ok := dict.Define(word)
		if !ok {
			fmt.Println("No definition found")
		} else {
			fmt.Println(definition)
		}
	}
}

func spellCheckerMenu() error {
	sc, err := spellchecker.NewEnglishSpellChecker()
	if err != nil {
		return err
	}
	for {
		fmt.Println()
		fmt.Print("Enter a sentence (type 0 to exit): ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		sentence := scanner.Text()
		if sentence == "0" {
			return nil
		}

		validWords, invalidWords := sc.CheckSentence(sentence)
		fmt.Println("Valid words:", validWords)
		fmt.Println("Invalid words:", invalidWords)
	}

}
