package helper

// DereferenceString returns the value of the pointer passed in.
func DereferenceString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// DereferenceInt returns the value of the pointer passed in.
func DereferenceInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
