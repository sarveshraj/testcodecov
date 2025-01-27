package mypackage

import "errors"

// ConvIntToStr returns the word for the given integer (1-20) or an error if out of range
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
	case 10:
		return "Ten", nil
	case 11:
		return "Eleven", nil
	case 12:
		return "Twelve", nil
	case 13:
		return "Thirteen", nil
	case 14:
		return "Fourteen", nil
	case 15:
		return "Fifteen", nil
	case 16:
		return "Sixteen", nil
	case 17:
		return "Seventeen", nil
	case 18:
		return "Eighteen", nil
	case 19:
		return "Nineteen", nil
	case 20:
		return "Twenty", nil
	default:
		return "", errors.New("number out of range")
	}
}
