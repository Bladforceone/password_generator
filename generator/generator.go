package generator

import (
	"crypto/rand"
	"math/big"
)

var (
	letterRunes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberRunes  = "0123456789"
	specialRunes = "!@#$%^&*()-_=+[]{}|;:,.<>/?"
)

func GeneratePassword(length int, useLetters, useNumbers, useSpecial bool) (string, error) {
	var charset string

	if useLetters {
		charset += letterRunes
	}

	if useNumbers {
		charset += numberRunes
	}

	if useSpecial {
		charset += specialRunes
	}

	password := make([]byte, length)
	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[num.Int64()]
	}
	return string(password), nil
}
