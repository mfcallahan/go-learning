package greetings

import (
	"fmt"
    "errors"
)

// CustomHello() returns a greeting for the named person.
func CustomHello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
        return "", errors.New("Param 'name' cannot be empty")
	}

    return fmt.Sprintf("Hi, %v!", name), nil
}
