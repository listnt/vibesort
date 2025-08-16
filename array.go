package vibes

import "fmt"

// Sorts an array in ascending order
func Sort[T any](arr []T) ([]T, error) {
	return callArray[T](fmt.Sprintf("Sort this array: %v", arr))
}
