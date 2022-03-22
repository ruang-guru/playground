package main

import (

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const wordLength = 5
// const maxGuess = 6

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
	words := []string{}


	return words
}


func main() {
	dictionary := getDictionaryWords()
	// TODO: answer here
	word:=dictionary[6]
fmt.Println(word)


}