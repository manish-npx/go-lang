package main

import (
	"errors"
	"fmt"

	"example.com/errors-demo/app"
	"example.com/errors-demo/tea"
)

func handleError(err error) {
	if err == nil {
		return
	}

	// If it’s an AppError, we can act by Code
	if appErr, ok := err.(*app.AppError); ok {
		switch appErr.Code {
		case "NO_TEA":
			fmt.Println("🫖 Please buy some tea leaves.")
		case "NO_WATER":
			fmt.Println("🚱 Please refill the kettle.")
		case "BROKEN_KETTLE":
			fmt.Println("🔥 Repair your kettle.")
		default:
			fmt.Println("Unknown issue:", appErr)
		}
		return
	}

	// fallback for any other generic errors
	fmt.Println("Unhandled error:", err)
}

/* custom error example */
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid %s: %s", e.Field, e.Message)
}

func CreateUser(name string, age int, adds []string) error {
	if name == "" {
		return &ValidationError{"name", "can not be empty"}
	}
	if age < 18 {
		return &ValidationError{"age", "must be 18 or older"}
	}
	if len(adds) < 2 {
		return &ValidationError{"address", "must contain city and state"}
	}
	if adds[0] == "" {
		return &ValidationError{"city", "cannot be empty"}
	}
	if adds[1] == "" {
		return &ValidationError{"state", "cannot be empty"}
	}

	return nil
}

func main() {

	// --- Custom validation example ---
	err := CreateUser("Manish", 20, []string{"", ""})
	var vErr *ValidationError
	if errors.As(err, &vErr) {
		fmt.Printf("Field: %s, Message: %s\n", vErr.Field, vErr.Message)
	} else if err != nil {
		fmt.Println("Some other error:", err)
		// if i pass blank return then its not goes to next makeTea and did not check that or execute that functions

	} else {
		fmt.Println("✅ User created successfully.")
	}

	fmt.Println("------------------------")

	// --- App-level error example ---
	msg, err1 := tea.MakeTea(false, false, false)
	if err1 != nil {
		handleError(err1)
		return
	}

	fmt.Println(msg)
}
