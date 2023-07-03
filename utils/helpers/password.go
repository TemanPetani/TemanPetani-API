package helpers

import (
	"errors"
	"unicode"
)

func ValidatePassword(s string) error {
	if len(s) < 8 {
		return errors.New("Password harus memiliki setidaknya 8 karakter")
	}

	var lower, upper, number, special bool
	for _, c := range s {
		switch {
		case unicode.IsLower(c):
			lower = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsNumber(c):
			number = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}

	if lower && upper && number && special {
		return nil
	}

	return errors.New("Password harus mengandung kombinasi huruf kapital dan huruf kecil, angka, dan karakter khusus")
}
