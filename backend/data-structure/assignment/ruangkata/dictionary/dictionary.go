package dictionary

type Dictionary interface {
	Define(word string) (string, bool)
}
