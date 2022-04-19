package spellchecker

type SpellChecker interface {
	CheckWord(word string) bool
	CheckSentence(sentence string) (validWords []string, invalidWords []string)
}
