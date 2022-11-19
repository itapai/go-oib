package oib

import "testing"

func TestGenerate(t *testing.T) {
	code := Generate()

	length := len(code)
	if length != 11 {
		t.Errorf("Expected %d, got %d", 11, length)
	}

	err := IsValid(code)
	valid := err == nil
	if !valid {
		t.Errorf("Expected %v, got %v", true, valid)
	}
}

func TestIsValid(t *testing.T) {
	cases := []struct {
		code string
		want bool
	}{
		{"10000000000", true},
		{"10000000001", false},
		{"", false},
	}

	for _, tc := range cases {
		err := IsValid(tc.code)
		got := err == nil
		if tc.want != got {
			t.Errorf("Expected %v, but got %v", tc.want, got)
		}
	}
}
