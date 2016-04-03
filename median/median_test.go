package median

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func medianTest(t *testing.T, a, b []string, want string) {
	m, err := Median(a, b)
	if err != nil {
		t.Fatal(err)
	}

	if m != want {
		t.Fatalf("got '%v', want '%v'", m, want)
	}
}

func TestMedianBothArraysEmpty(t *testing.T) {
	m, err := Median(nil, nil)
	if err == nil {
		t.Fatal("expected error when both arrays are empty")
	}
	if m != "" {
		t.Fatalf("got '%v', want empty string", m)
	}
}

func TestMedianOneValue1(t *testing.T) {
	a := []string{"5"}
	var b []string
	medianTest(t, a, b, "5")
}

func TestMedianOneValue2(t *testing.T) {
	var a []string
	b := []string{"5"}
	medianTest(t, a, b, "5")
}

func TestMedianFirstEmpty(t *testing.T) {
	var a []string
	b := []string{"2", "4", "6", "8", "9"}
	medianTest(t, a, b, "6")
}

func TestMedianSecondEmpty(t *testing.T) {
	a := []string{"2", "4", "6", "8", "9"}
	var b []string
	medianTest(t, a, b, "6")
}

func TestMedianSingleElements(t *testing.T) {
	a := []string{"5"}
	b := []string{"6"}
	medianTest(t, a, b, "5")
}

func TestMedianOddOdd(t *testing.T) {
	a := []string{"01", "03", "05", "07", "09"}
	b := []string{"02", "04", "06", "08", "10"}
	medianTest(t, a, b, "05")
}

func TestMedianEvenEven(t *testing.T) {
	a := []string{"00", "01", "03", "05", "07", "09"}
	b := []string{"02", "04", "06", "08", "10", "11"}
	medianTest(t, a, b, "05")
}

func TestMedianOddEven(t *testing.T) {
	a := []string{"1", "3", "5", "7", "9"}
	b := []string{"2", "4", "6", "8"}
	medianTest(t, a, b, "5")
}

func TestMedianEvenOdd(t *testing.T) {
	a := []string{"2", "4", "6", "8"}
	b := []string{"1", "3", "5", "7", "9"}
	medianTest(t, a, b, "5")
}

func TestMedianExact(t *testing.T) {
	a := []string{"1", "3", "5", "7", "9"}
	b := []string{"2", "4", "5", "6", "8"}
	medianTest(t, a, b, "5")
}

func TestMedian1on2Case1(t *testing.T) {
	a := []string{"1"}
	b := []string{"2", "3"}
	medianTest(t, a, b, "2")
}

func TestMedian1on2Case2(t *testing.T) {
	a := []string{"2"}
	b := []string{"1", "3"}
	medianTest(t, a, b, "2")
}

func TestMedian1on2Case3(t *testing.T) {
	a := []string{"3"}
	b := []string{"1", "2"}
	medianTest(t, a, b, "2")
}

func TestMedian2on1Case1(t *testing.T) {
	a := []string{"2", "3"}
	b := []string{"1"}
	medianTest(t, a, b, "2")
}

func TestMedian2on1Case2(t *testing.T) {
	a := []string{"1", "3"}
	b := []string{"2"}
	medianTest(t, a, b, "2")
}

func TestMedian2on1Case3(t *testing.T) {
	a := []string{"1", "2"}
	b := []string{"3"}
	medianTest(t, a, b, "2")
}

func TestMedian2on2Case1(t *testing.T) {
	a := []string{"1", "2"}
	b := []string{"3", "4"}
	medianTest(t, a, b, "2")
}

func TestMedian2on2Case2(t *testing.T) {
	a := []string{"3", "4"}
	b := []string{"1", "2"}
	medianTest(t, a, b, "2")
}

func TestMedian2on2Case3(t *testing.T) {
	a := []string{"1", "3"}
	b := []string{"2", "4"}
	medianTest(t, a, b, "2")
}

func TestMedian2on2Case4(t *testing.T) {
	a := []string{"2", "4"}
	b := []string{"1", "3"}
	medianTest(t, a, b, "2")
}

func TestMedian2on2Case5(t *testing.T) {
	a := []string{"1", "4"}
	b := []string{"2", "3"}
	medianTest(t, a, b, "2")
}

func TestMedian2on2Case6(t *testing.T) {
	a := []string{"2", "3"}
	b := []string{"1", "4"}
	medianTest(t, a, b, "2")
}

func TestMedianOn10Case1(t *testing.T) {
	a := []string{"01", "02", "03", "06", "07", "08", "09"}
	b := []string{"04", "05", "10"}
	medianTest(t, a, b, "05")
}

func TestMedianOn10Case2(t *testing.T) {
	a := []string{"04", "05", "10"}
	b := []string{"01", "02", "03", "06", "07", "08", "09"}
	medianTest(t, a, b, "05")
}

func TestMedianOn10Case3(t *testing.T) {
	a := []string{"1", "2", "3", "5", "8", "9"}
	b := []string{"4", "6", "7"}
	medianTest(t, a, b, "5")
}

func TestMedianOn10Case4(t *testing.T) {
	a := []string{"6", "7"}
	b := []string{"1", "2", "3", "4", "5", "8", "9"}
	medianTest(t, a, b, "5")
}

func TestMedianOn10Case5(t *testing.T) {
	a := []string{"3", "9"}
	b := []string{"1", "2", "4", "5", "6", "7", "8"}
	medianTest(t, a, b, "5")
}

func TestMedianOn30Case1(t *testing.T) {
	a := []string{"01", "02", "05", "06", "08", "09", "12", "13", "15", "16", "17", "18", "21", "22", "23"}
	b := []string{"03", "04", "07", "10", "11", "14", "19", "20", "24", "25", "26", "27", "28", "29", "30"}
	medianTest(t, a, b, "15")
}

func TestMedianOn100Case1(t *testing.T) {
	a := []string{"01", "02", "03", "04", "05", "07", "08", "09", "11", "12", "13", "16", "18", "25", "27", "28", "30", "32", "35", "36", "38", "40", "42", "44", "48", "49", "50", "56", "57", "60", "61", "66", "69", "70", "71", "73", "74", "75", "76", "78", "80", "82", "85", "86", "90", "92", "93", "97", "98", "99"}
	b := []string{"06", "10", "14", "15", "17", "19", "20", "21", "22", "23", "24", "26", "29", "31", "33", "34", "37", "39", "41", "43", "45", "46", "47", "51", "52", "53", "54", "55", "58", "59", "62", "63", "64", "65", "67", "68", "72", "77", "79", "81", "83", "84", "87", "88", "89", "91", "94", "95", "96"}
	medianTest(t, a, b, "50")
}

func TestMedianOn100Case2(t *testing.T) {
	a := []string{"03", "08", "22", "23", "25", "26", "29", "30", "36", "37", "42", "47", "50", "53", "56", "61", "62", "64", "65", "66", "68", "73", "74", "75", "76", "77", "81", "84", "87", "88", "89", "92", "93", "95", "97", "98", "99"}
	b := []string{"01", "02", "04", "05", "06", "07", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "24", "27", "28", "31", "32", "33", "34", "35", "38", "39", "40", "41", "43", "44", "45", "46", "48", "49", "51", "52", "54", "55", "57", "58", "59", "60", "63", "67", "69", "70", "71", "72", "78", "79", "80", "82", "83", "85", "86", "90", "91", "94", "96"}
	medianTest(t, a, b, "50")
}

func TestMedianOn100Case3(t *testing.T) {
	a := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "70", "71", "72", "73", "74", "75", "76", "77", "78", "79", "80", "81", "82", "83", "84", "85", "86", "87", "88", "89", "90", "91", "92", "93", "94", "95", "96", "97", "98", "99"}
	b := []string{"47"}
	medianTest(t, a, b, "50")
}

// BenchmarkMedian creates a valid random test every time it gets called.
// It allows s and t to have any size each.
func BenchmarkMedian(b *testing.B) {
	const size = 1000 // power of 10
	const want = size / 2

	s := make([]string, 0, size)
	t := make([]string, 0, size)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	pivot := r.Float64()

	// Make two slices, s and t, such that both are sorted and
	// "s MERGE t = [1..size-1], but each has a random subset of [1..size-1]
	for i := 1; i < size; i++ {
		if r.Float64() < pivot {
			s = append(s, fmt.Sprintf("%03d", i))
		} else {
			t = append(t, fmt.Sprintf("%03d", i))
		}
	}

	swant := fmt.Sprintf("%03d", want)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m, err := Median(s, t)
		if err != nil {
			b.Fatal(err)
		}

		if m != swant {
			b.Fatalf("\na := %#v\nb := %#v\nmedianTest(t, a, b, \"%v\")\ngot '%v'", s, t, swant, m)
		}
	}
}

// BenchmarkMedian creates a valid random test every time it gets called.
// It forces s and t to have the same size.
func BenchmarkMedianEqualSizes(b *testing.B) {
	const size = 1000
	const want = size / 2

	s := make([]string, 0, size/2)
	t := make([]string, 0, size/2)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	pivot := r.Float64()

	// Make two slices, s and t, such that both are sorted and
	// "s MERGE t = [1..size-1], but each has a random subset of [1..size-1]
	for i := 1; i <= size; i++ {
		var tos bool
		switch {
		case len(t) == (size / 2):
			tos = true
		case len(s) == (size / 2):
			tos = false
		default:
			tos = (r.Float64() < pivot)
		}
		if tos {
			s = append(s, fmt.Sprintf("%03d", i))
		} else {
			t = append(t, fmt.Sprintf("%03d", i))
		}
	}

	swant := fmt.Sprintf("%03d", want)

	if len(s) != len(t) {
		b.Fatalf("creation of random test failed: len(s)=%v != %v=len(t)", len(s), len(t))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m, err := Median(s, t)
		if err != nil {
			b.Fatal(err)
		}

		if m != swant {
			b.Fatalf("\na := %#v\nb := %#v\nmedianTest(t, a, b, \"%v\")\ngot '%v'", s, t, swant, m)
		}
	}
}
