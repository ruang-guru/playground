package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

const wordLength = 5
const maxGuess = 6

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
					//when the answer is:
					//STROK, and we guess:
					//SOSOK
					//the answer should be:
					//GXXGG
					//not:
					//GYYGG
					//Reason: the second 'O' has been marked as correct position ('Y')
					//if we mark 'Y' for the first 'O', people would guess there should be yet another 'O'
					//while in fact there is only one 'O' in 'STROK'
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
	dictionary := getDictionaryWords()

	// TODO: answer here
}
