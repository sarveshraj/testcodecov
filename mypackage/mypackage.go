package mypackage

import "errors"

// ConvIntToStr returns the word for the given integer (1-9) or an error if out of range
func ConvIntToStr(n int) (string, error) {
	switch n {
	case 1:
		return "One", nil
	case 2:
		return "Two", nil
	case 3:
		return "Three", nil
	case 4:
		return "Four", nil
	case 5:
		return "Five", nil
	case 6:
		return "Six", nil
	case 7:
		return "Seven", nil
	case 8:
		return "Eight", nil
	case 9:
		return "Nine", nil
	default:
		return "", errors.New("number out of range")
	}
}
