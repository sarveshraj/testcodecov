package mypackage

import "testing"

func TestConvIntToStr(t *testing.T) {
	tests := []struct {
		input    int
		expected string
		err      bool
	}{
		{1, "One", false},
		{5, "Five", false},
		{8, "Eight", false},
		{10, "", true},
	}

	for _, test := range tests {
		result, err := ConvIntToStr(test.input)
		if (err != nil) != test.err {
			t.Errorf("ConvIntToStr(%d) error = %v, wantErr %v", test.input, err, test.err)
			continue
		}
		if result != test.expected {
			t.Errorf("ConvIntToStr(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}
