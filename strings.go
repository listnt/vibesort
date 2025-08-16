package vibes

import "fmt"

// To lowercase
func ToLower(s string) (string, error) {
	return call[string](fmt.Sprintf("To lower: %s", s))
}

// To uppercase
func ToUpper(s string) (string, error) {
	return call[string](fmt.Sprintf("To upper: %s", s))
}
