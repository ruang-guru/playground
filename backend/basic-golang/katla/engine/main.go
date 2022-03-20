package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const wordLength = 5
const maxGuess = 6

//NOTE: err handling is not yet taught, we don't handle errors in this example
//don't worry about the content of this method for now. We haven't learn some concepts
func getDictionaryWords() []string {
	resp, err := http.Get("https://kbbi.vercel.app")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var jsonResp struct {
		Entries []string `json:"entries"` //again, don't worry about this for now
	}
	json.Unmarshal(body, &jsonResp)
	words := make([]string, 0)

	for _, entry := range jsonResp.Entries {
		tokens := strings.Split(entry, "/")
		word := tokens[len(tokens)-1]
		matched, _ := regexp.MatchString(fmt.Sprintf("^[a-zA-Z]{%d}$", wordLength), word) //don't worry about this
		if matched {
			words = append(words, strings.ToLower(word))
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

	rand.Seed(time.Now().UnixNano())
	answer := dictionary[rand.Intn(len(dictionary))]
	answer = "strok"
	// fmt.Printf("Answer: %s\n", answer)

	isWin := false
	for trial := 0; trial < maxGuess; trial++ {
		var guess string
		fmt.Printf("Guess %d: \n", trial+1)
		for {
			guess = ""
			fmt.Scanln(&guess)

			if len(guess) != wordLength {
				fmt.Printf("Please enter exactly %d characters\n", len(answer))
				continue
			}

			isAllLowerCase := true
			for _, c := range guess {
				if !(c >= 'a' && c <= 'z') {
					isAllLowerCase = false
				}
			}
			if !isAllLowerCase {
				fmt.Println("Please enter lowercase characters only")
				continue
			}

			if !isInDictionary(guess, dictionary) {
				fmt.Printf("%s is not in the dictionary\n", guess)
				continue
			}

			break
		}

		hints := calculateHints(guess, answer)
		for i := 0; i < wordLength; i++ {
			if hints[i] == notFound {
				fmt.Printf("X")
			} else if hints[i] == correctPosition {
				fmt.Printf("G")
			} else if hints[i] == correctLetter {
				fmt.Printf("Y")
			}
		}
		fmt.Println()
		fmt.Println()

		if guess == answer {
			fmt.Println("You win!")
			isWin = true
			break
		}
	}

	if !isWin {
		fmt.Printf("The correct answer is: %s\n", answer)
	}
}
