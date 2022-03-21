package answerremover

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func removeTrailingSpace(s string) (withoutTrailingSpace string, trailingSpace string) {
	m := regexp.MustCompile(`(^\s+)`)
	withoutTrailingSpace = m.ReplaceAllString(s, "")
	trailingSpaces := m.FindStringSubmatch(s)

	if len(trailingSpaces) > 0 {
		trailingSpace = trailingSpaces[1]
	} else {
		trailingSpace = ""
	}
	return
}

func RemoveAnswerBlock(text string) (string, error) {
	lines := strings.Split(text, "\n")
	isAnswerBlock := false
	res := NewAccumulator()

	for _, line := range lines {
		trimmedLine, trailingSpace := removeTrailingSpace(line)
		if match := matchBeginAnswer(trimmedLine); match != nil {
			if isAnswerBlock {
				return "", errors.New("beginanswer found twice")
			}
			isAnswerBlock = true
		} else {
			endAnswer := matchEndAnswer(trimmedLine)
			if endAnswer != nil {
				if !isAnswerBlock {
					return "", errors.New("found endanswer but no matched beginanswer")
				} else {
					processor := NewProcessor(endAnswer.annotation)
					replacedLine, err := processor.Process(trailingSpace, res.getEnclosedLines(), endAnswer.variation)

					if err != nil {
						return "", err
					}
					if replacedLine != nil {
						res.addOutputLine(*replacedLine)
					}
				}
				isAnswerBlock = false
			} else if isAnswerBlock {
				res.addEnclosedLine(line)
			} else {
				res.addOutputLine(line)
			}
		}
	}
	if isAnswerBlock {
		return "", errors.New("answer block is not closed")
	}
	return strings.Join(res.getOutputs(), "\n"), nil
}

func RemoveAllAnswerBlocks(rootFolder string, processedFileNames []string, processedExtensions []string, excludeFolders []string) error {
	return filepath.WalkDir(rootFolder,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}

			{
				processed := false
				for _, ext := range processedExtensions {
					if strings.HasSuffix(path, ext) {
						processed = true
						break
					}
				}
				for _, name := range processedFileNames {
					if d.Name() == name {
						processed = true
						break
					}
				}
				if !processed {
					return nil
				}
			}

			for _, excludeFolder := range excludeFolders {
				if strings.HasPrefix(path, excludeFolder) {
					return nil
				}
			}

			fmt.Print("Processing: ", path)

			text, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			newText, err := RemoveAnswerBlock(string(text))
			if err != nil {
				fmt.Println(" ❌")
				return err
			} else {
				if string(text) != newText {
					fmt.Println(" ✅")
				} else {
					fmt.Println(" ➖")
				}
			}
			return os.WriteFile(path, []byte(newText), 0644)
		},
	)
}
