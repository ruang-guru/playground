package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

const wordLength = 6
const maxGuess = 8

//NOTE: err handling is not yet taught, we don't handle errors in this example
//don't worry about the content of this method for now. We haven't learn some concepts
func getDictionaryWords() []string {
	kbbiURL := "https://gist.githubusercontent.com/fikriauliya/c7024f9629ba7d515f01a625c66a4f2f/raw/141f629d452145ce7e02215a98cde04d9f1bbb20/kbbi.txt"

	data, err := http.Get(kbbiURL)
	if err != nil {
		panic(err)
	}
	defer data.Body.Close()

	body, err := io.ReadAll(data.Body)
	if err != nil {
		panic(err)
	}

	lines := string(body)
	words := make([]string, 0)

	for _, line := range strings.Split(lines, "\n") {
		matched, _ := regexp.MatchString(fmt.Sprintf("^[a-zA-Z]{%d}$", wordLength), line) //don't worry about this
		if matched {
			words = append(words, strings.ToLower(line))
		}
	}
	return words
}

// this is not optimal search algorithm, but it's a good example
func isInDictionary(word string, dictionary []string) bool {
	for _, entry := range dictionary {
		if entry == word {
			return true
		}
	}
	return false
}

type hint int

const (
	notFound        hint = 0
	correctPosition hint = 1
	correctLetter   hint = 2
)

func calculateHints(guess, answer string) (hints []hint) {
	guessChars := []rune(guess)
	answerChars := []rune(answer)

	hints = make([]hint, wordLength)

	for i := 0; i < wordLength; i++ {
		if guessChars[i] == answerChars[i] {
			hints[i] = correctPosition
		} else {
			for j := 0; j < wordLength; j++ {
				if i != j {
					if guessChars[i] == answerChars[j] && guessChars[j] != answerChars[j] {
						hints[i] = correctLetter
						break
					}
				}
			}
		}
	}
	return
}

func main() {
	isWin := false
	dictionary := getDictionaryWords()

<<<<<<< HEAD
	// TODO: answer here
	word:=dictionary[rand.Intn(len(dictionary))]
	// fmt.Println(word)

	for gameCount := 1; gameCount <= maxGuess; gameCount++ {
		var guess string
		
		fmt.Println("Attempt number ",gameCount," :")
		for{		
			fmt.Scanln(&guess)
			if len(guess)<wordLength || len(guess)>wordLength{
				fmt.Println("word length is invalid")
				fmt.Println("Attempt number ",gameCount," :")
				continue
			}
			guess = strings.ToLower(guess)

			if !isInDictionary(guess,dictionary) {
				fmt.Println("the answer is not in dictionary")
				fmt.Println("Attempt number ",gameCount," :")
				continue
			}
			break

		}
		hints:= calculateHints(guess,word)

		for i := 0; i < wordLength; i++ {
			if hints[i]==correctLetter {
				fmt.Print("Y")
				
			}else if hints[i]==correctPosition {
				fmt.Print("G")
				
			}else if hints[i]==notFound {
				fmt.Print("X")
				
			}
		}
		for i := 0; i < 3; i++ {
			fmt.Println()
		}

		if guess==word {
			fmt.Println("you win")
			isWin =true
			break;
		}

	}

	if !isWin{
		fmt.Println("You Lost, the answer is : ", word)
	}


=======
	//beginanswer
	for len(dictionary) > 1 {
		var guess, coloredHints string
		fmt.Printf("Guess: ")

		fmt.Scanln(&guess)

		fmt.Printf("Hint: ")
		fmt.Scanln(&coloredHints)

		receivedHints := make([]hint, wordLength)
		for i := 0; i < wordLength; i++ {
			switch coloredHints[i] {
			case 'X':
				receivedHints[i] = notFound
			case 'Y':
				receivedHints[i] = correctLetter
			case 'G':
				receivedHints[i] = correctPosition
			}
		}

		filteredDictionary := make([]string, 0)

		for _, dict := range dictionary {
			hints := calculateHints(guess, dict)
			match := true

			for i := 0; i < wordLength; i++ {
				if hints[i] != receivedHints[i] {
					match = false
					break
				}
			}

			if match {
				filteredDictionary = append(filteredDictionary, dict)
			}
		}
		dictionary = filteredDictionary

		fmt.Print("Possible words: ")
		for _, dict := range dictionary {
			fmt.Printf("%s ", dict)
		}

		fmt.Println()
		fmt.Println()
	}
	//endanswer
>>>>>>> 61b24143472cbe92924703fbc80d83dff954219b
}
