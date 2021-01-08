package main

import (
	"fmt"
	"log"
)

import "rsc.io/quote"
import "example.com/greetings"

func main() {
	fmt.Println(quote.Go())

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	//log.SetPrefix("greetings: ")
	//log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("123")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
			log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	greetings.Trace(message)

    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}
    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    // If no error was returned, print the returned map of
    // messages to the console.
    greetings.Trace("%v", messages)

}

