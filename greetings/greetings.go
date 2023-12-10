package greetins

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("Empty name")
	}
	// Return a greeting that embeds the name in a message
	// if a name was recived
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// returns one of a set of greeting messages. The returned
// message is selected at random
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	// return a randomly selected messafe format by specifyng
	// a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
