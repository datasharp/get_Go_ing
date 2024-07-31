package greetings

import (
	"fmt"
	"errors"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// error handling
	if name == "" {
		return "", errors.New("empty name")
	}
	
	// return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func Goodbye(name string) string {
	message := fmt.Sprintf("Goodbye, %v. Talk soon!", name)
	return message
}



// := -> value on right as value for message
// var message string
// message = fmt.Sprintf("Hi, %v. Welcome!", name)
