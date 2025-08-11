package vibesort

import (
	"reflect"
	"testing"
)

func TestInts(t *testing.T) {
	v := NewVibesort("your-test-key")

	input := []int64{42, 7, 88, 15, 63}
	expected := []int64{7, 15, 42, 63, 88}

	res, err := v.Vibesort(input)
	t.Log(res, err)

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected %v, got %v", expected, res)
	}
}
