package rpn

import (
	"testing"
)

func TestEvalrpnNoTokens(t *testing.T) {
	_, err := Evalrpn(nil)
	if err == nil {
		t.Fatal("Expected error for empty token list.")
	}
}

func TestEvalrpn1(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}

	result, err := Evalrpn(tokens)
	if err != nil {
		t.Fatalf("unexpected error: %v.", err)
	}

	if result != 9 {
		t.Fatalf("for %v got %v, wanted %v", tokens, result, 9)
	}
}

func TestEvalrpn2(t *testing.T) {
	tokens := []string{"4", "13", "5", "/", "+", "10", "*"}

	result, err := Evalrpn(tokens)
	if err != nil {
		t.Fatalf("unexpected error: %v.", err)
	}

	if result != 66 {
		t.Fatalf("for %v got %v, wanted %v", tokens, result, 66)
	}
}

func TestEvalrpn3(t *testing.T) {
	tokens := []string{"2", "3", "^", "1", "-"}

	result, err := Evalrpn(tokens)
	if err != nil {
		t.Fatalf("unexpected error: %v.", err)
	}

	if result != 7 {
		t.Fatalf("for %v got %v, wanted %v", tokens, result, 7)
	}
}

func TestEvalrpnSumNoOperands(t *testing.T) {
	tokens := []string{"+"}

	_, err := Evalrpn(tokens)
	if err == nil {
		t.Fatal("Expected error for sum without operands.")
	}
}

func TestEvalrpnSumOneOperand(t *testing.T) {
	tokens := []string{"1", "+"}

	_, err := Evalrpn(tokens)
	if err == nil {
		t.Fatal("Expected error for sum with operand.")
	}
}

func TestEvalrpnBadTokeb(t *testing.T) {
	tokens := []string{"1", "hola", "+"}

	_, err := Evalrpn(tokens)
	if err == nil {
		t.Fatal("Expected error for unexpected token.")
	}
}

func BenchmarkEvalrpn(b *testing.B) {
	tokens := []string{"4", "13", "5", "/", "+", "10", "*"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Evalrpn(tokens)
	}
}
