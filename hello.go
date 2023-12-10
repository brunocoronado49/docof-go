package main

import (
	"fmt"
	"log"

	greetins "github.com/brunocoronado49/docof-go/greetings"
)

func main() {
	// Set the properties of the predefined logger, including
	// the log entry prefix and a flag to disable prirting
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// request a greeting message
	// if an error was returned, print it to the console and
	// exit the program
	message, err := greetins.Hello("Gladys")
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned message
	// to the console
	fmt.Println(message)
}
