package vibes

import (
	"cmp"
	"fmt"
)

// Adds two entities
func Add[T cmp.Ordered](a T, b T) (T, error) {
	return call[T](fmt.Sprintf("Add these numbers: %v %v", a, b))
}

// Substract two entities
func Sub[T cmp.Ordered](a T, b T) (T, error) {
	return call[T](fmt.Sprintf("Substract second from first: %v %v", a, b))
}

// Multiplies two entities
func Mult[T cmp.Ordered](a T, b T) (T, error) {
	return call[T](fmt.Sprintf("Multiply these numbert: %v %v", a, b))
}

// Divides two entities
func Div[T cmp.Ordered](a T, b T) (T, error) {
	return call[T](fmt.Sprintf("Divide first by second: %v %v", a, b))
}
