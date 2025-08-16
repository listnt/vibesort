package vibes

import (
	"reflect"
	"testing"
)

// Seeks 29th Fibonacci number
func TestFibonacci(t *testing.T) {
	InitVibes("your-key-here")

	fibNumbers := []int64{1, 1}
	for i := 0; i < 29-2; i++ {
		newFib, err := Add(fibNumbers[len(fibNumbers)-1], fibNumbers[len(fibNumbers)-2])
		if err != nil {
			t.Fatal(err)
		}

		fibNumbers = append(fibNumbers, newFib)
	}

	if fibNumbers[len(fibNumbers)-1] != 514229 {
		t.Fatal("failed to calc fib sequence")
	}
}

func TestSort(t *testing.T) {
	InitVibes("your-key-here")

	input := []int64{42, 7, 88, 15, 63}
	expected := []int64{7, 15, 42, 63, 88}

	res, err := Sort(input)
	t.Log(res, err)

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected %v, got %v", expected, res)
	}
}
