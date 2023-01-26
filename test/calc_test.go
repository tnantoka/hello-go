package calc

import "testing"

func TestAdd(t *testing.T) {
	a, b, expected := 1, 2, 3
	acctual := Add(a, b)
	if acctual != expected {
		t.Errorf("Add(%d, %d) = %d; expected %d", a, b, acctual, expected)
	}
}

func TestSub(t *testing.T) {
	var tests = []struct {
		a, b, expected int
	}{
		{1, 2, -1},
		{2, 1, 1},
		{0, 0, 0},
	}
	for i, test := range tests {
		actual := Sub(test.a, test.b)
		if actual != test.expected {
			t.Errorf("Test %v: Sub(%v, %v) = %v; expected %v", i, test.a, test.b, actual, test.expected)
		}
	}
}

// Test Div with table-driven tests
func TestDiv(t *testing.T) {
	var tests = []struct {
		a, b, expected int
	}{
		{1, 2, 0},
		{2, 1, 2},
		{0, 1, 0},
		{1, 0, 0},
	}
	for i, test := range tests {
		actual := Div(test.a, test.b)
		if actual != test.expected {
			t.Errorf("Test %v: Div(%v, %v) = %v; expected %v", i, test.a, test.b, actual, test.expected)
		}
	}
}
