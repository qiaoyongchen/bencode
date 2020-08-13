package token

import "strconv"

// IsDigit IsDigit
func IsDigit(s string) bool {
	_, ierr := strconv.Atoi(s)
	if ierr != nil {
		return false
	}
	return true
}
