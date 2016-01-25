package rpn

import (
	"testing"
)

func TestEvalrpn(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}

	result, err := Evalrpn(tokens)
	if err != nil {
		t.Fatalf("unexpected error: %v.", err)
	}

	if result != 9 {
		t.Fatalf("for %v got %v, wanted %v", tokens, result, 9)
	}

	tokens = []string{"4", "13", "5", "/", "+", "10", "*"}

	result, err = Evalrpn(tokens)
	if err != nil {
		t.Fatalf("unexpected error: %v.", err)
	}

	if result != 66 {
		t.Fatalf("for %v got %v, wanted %v", tokens, result, 66)
	}

	tokens = []string{"2", "3", "^"}

	result, err = Evalrpn(tokens)
	if err != nil {
		t.Fatalf("unexpected error: %v.", err)
	}

	if result != 8 {
		t.Fatalf("for %v got %v, wanted %v", tokens, result, 8)
	}
}
