package utils

import (
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePasswords(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func PasswordStrength(password string) bool {
	strength := 0

	if len(password) >= 7 {
		strength += 2
	}

	hasUppercase := false
	hasLowercase := false
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUppercase = true
		}
		if unicode.IsLower(char) {
			hasLowercase = true
		}
	}
	if hasUppercase && hasLowercase {
		strength += 2
	}

	hasDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}
	if hasDigit {
		strength++
	}

	hasSpecial := false
	for _, char := range password {
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			hasSpecial = true
			break
		}
	}
	if hasSpecial {
		strength++
	}

	return strength >= 5
}

func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regex := regexp.MustCompile(emailRegex)

	return regex.MatchString(email)
}
