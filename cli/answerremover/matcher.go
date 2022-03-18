package answerremover

import (
	"regexp"
	"strings"
)

func matchBeginAnswer(line string) *commentvariation {
	variations := []*regexp.Regexp{
		regexp.MustCompile(`^//\s*beginanswer\s*$`),
		regexp.MustCompile(`^--\s*beginanswer\s*$`),
		regexp.MustCompile(`^#\s*beginanswer\s*$`),
		regexp.MustCompile(`^/\*\s*beginanswer\s*\*/\s*$`),
		regexp.MustCompile(`^{\s*/\*\s*beginanswer\s*\*/\s*}\s*$`),
		regexp.MustCompile(`<!--\s*beginanswer\s*-->$`),
	}
	for i, m := range variations {
		if m.MatchString(line) {
			variation := commentvariation(i)
			return &variation
		}
	}
	return nil
}

func isOnlySpace(s string) bool {
	return len(s) == 0 || strings.TrimSpace(s) == ""
}

func matchEndAnswer(line string) *endanswer {
	variations := []*regexp.Regexp{
		regexp.MustCompile(`^//\s*endanswer(\s*(.+?))?\s*$`),
		regexp.MustCompile(`^--\s*endanswer(\s*(.+?))?\s*$`),
		regexp.MustCompile(`^#\s*endanswer(\s*(.+?))?\s*$`),
		regexp.MustCompile(`^/\*\s*endanswer(\s*(.+?))?\s*\*/\s*$`),
		regexp.MustCompile(`^{\s*/\*\s*endanswer(\s*(.+?))?\*/\s*}\s*$`),
		regexp.MustCompile(`<!--\s*endanswer(\s*(.+?))?\s*-->$`),
	}
	for i, m := range variations {
		if m.MatchString(line) {
			match := m.FindStringSubmatch(line)
			variation := commentvariation(i)
			if len(match) > 1 {
				annotation := &match[2]
				if isOnlySpace(*annotation) {
					annotation = nil
				}
				return &endanswer{annotation, &variation}
			}

			return &endanswer{nil, &variation}
		}
	}
	return nil
}
