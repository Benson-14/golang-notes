package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   *int   `json:"age"` // Pointer to distinguish nil (not provided) from 0
}

type ValidationError struct {
	Field   string
	Message string
}

func ParseJSON(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}

func ValidateUser(user User) []ValidationError {
	// Return validation errors for:
	// - Name: required, 2-50 characters
	// - Email: required, must contain "@"
	// - Age: if not nil, must be 1-150
	var errors []ValidationError

	if user.Name == "" {
		errors = append(errors, ValidationError{"name", "is required"})
	} else if len(user.Name) < 2 || len(user.Name) > 50 {
		errors = append(errors, ValidationError{"name", "must be between 2 and 50 characters"})
	}

	if user.Email == "" {
		errors = append(errors, ValidationError{"email", "is required"})
	} else if !strings.Contains(user.Email, "@") {
		errors = append(errors, ValidationError{"email", "must contain '@'"})
	}

	if user.Age != nil {
		if *user.Age < 1 || *user.Age > 150 {
			errors = append(errors, ValidationError{"age", "must be between 1-15-"})
		}
	}

	return errors
}

func main() {
	jsonBody := strings.NewReader(`{"name": "Jo", "email": "invalid", "age": 200}`)

	var user User
	if err := ParseJSON(jsonBody, &user); err != nil {
		fmt.Printf("Parse error: %v\n", err)
		return
	}

	fmt.Printf("Parsed: %+v\n", user)

	errors := ValidateUser(user)
	fmt.Printf("Validation errors: %d\n", len(errors))
	for _, e := range errors {
		fmt.Printf("  %s: %s\n", e.Field, e.Message)
	}
}