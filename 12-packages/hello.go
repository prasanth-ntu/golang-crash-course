/*
Sources used for this example:

Create a go module: https://go.dev/doc/tutorial/create-module
Call your code from another module: https://go.dev/doc/tutorial/call-module-code
Return and handle an error: https://go.dev/doc/tutorial/handle-errors
Return a random greeting: https://go.dev/doc/tutorial/random-greeting
Return greetings for multiple people: https://go.dev/doc/tutorial/greetings-multiple-people
Add a test: https://go.dev/doc/tutorial/add-a-test

*/

package main

import (
	"fmt"
	"log"

	"greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)

	// Iterate over map items //
	for key, value := range messages {
		fmt.Println("Name:", key, "=>", "Message:", value)
	}
}
