package utils

import "fmt"

var (
	ErrInvalidInput = func(input string) error {
		return fmt.Errorf("invalid input: %s", input)
	}

	ErrNotFound = func(resource string, id int) error {
		return fmt.Errorf("%s with ID %d not found", resource, id)
	}
	ErrNotFoundUrl = func(resource string, url string) error {
		return fmt.Errorf("%s with Url %s not found", resource, url)
	}
	ErrNotFoundUser = func(resource string, email string) error {
		return fmt.Errorf("%s with email %s not found", resource, email)
	}
)
