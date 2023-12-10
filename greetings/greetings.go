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

// Returns a map that associates each of the named people
// witch a greeting message
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with message
	messages := make(map[string]string)

	// Loop through the recived slice of names, calling
	// the hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map associate the retrived message
		// with the name
		messages[name] = message
	}

	return messages, nil
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
