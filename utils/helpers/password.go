package helpers

import (
	"errors"
	"unicode"
)

func ValidatePassword(s string) error {
	var lower, upper, number, special bool
	chars := 0

	for _, c := range s {
		switch {
		case unicode.IsLower(c):
			lower = true
			chars++
		case unicode.IsUpper(c):
			upper = true
			chars++
		case unicode.IsNumber(c):
			number = true
			chars++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
			chars++
		}
	}

	if chars <= 8 {
		return errors.New("Password harus memiliki setidaknya 8 karakter")
	}

	if !lower && !upper && !number && !special {
		return errors.New("Password harus mengandung kombinasi huruf kapital dan huruf kecil, angka, dan karakter khusus")
	}

	return nil
}
