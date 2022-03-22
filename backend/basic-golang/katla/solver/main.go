package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const wordLength = 6
const maxGuess = 8

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


}
