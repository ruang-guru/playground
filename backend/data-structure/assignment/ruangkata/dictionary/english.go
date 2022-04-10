package dictionary

import (
	"bufio"
	"encoding/csv"
	"os"
	"strings"

	"github.com/ruang-guru/playground/backend/data-structure/assignment/ruangkata/downloader"
)

type dictionary struct {
	wordDefinition map[string]string
}

const url string = "https://gist.githubusercontent.com/fikriauliya/fcf4ef069f06550d4c3792b6cbd83abc/raw/3e1a1776c5f8dd871c43c5727c5cb9b1d1b0ab78/english.csv"
const filePath string = "./english-definition.csv"

func parseCSV() (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records, err := csv.NewReader(bufio.NewReader(file)).ReadAll()
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, record := range records {
		word := record[0]
		word = strings.ToLower(word)
		definition := record[2]
		result[word] = definition
	}
	return result, nil
}

func NewEnglishDictionary() (Dictionary, error) {
	err := downloader.Download(url, filePath)
	if err != nil {
		return nil, err
	}
	wordDefinition, err := parseCSV()
	if err != nil {
		return nil, err
	}
	return &dictionary{wordDefinition}, nil
}

func (d *dictionary) Define(word string) (string, bool) {
	word = strings.ToLower(word)
	if definition, ok := d.wordDefinition[word]; ok {
		return definition, true
	}
	return "", false
}
