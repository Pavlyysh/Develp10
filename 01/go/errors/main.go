package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Value   any
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: field %s with value %v - %s", e.Field, e.Value, e.Message)
}

func main() {
	err := ValidateAge(12)
	if err != nil {
		var valErr *ValidationError
		if errors.As(err, &valErr) {
			fmt.Println("custom err")
			fmt.Printf("Field: %s\nValue: %v\nMessage: %s\n", valErr.Field, valErr.Value, valErr.Message)
		}
	}
}

func ValidateAge(age int) error {
	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Value:   age,
			Message: "age cannot be negative",
		}
	}
	if age < 18 {
		return &ValidationError{
			Field:   "age",
			Value:   age,
			Message: "age is under 18",
		}
	}

	return nil
}
