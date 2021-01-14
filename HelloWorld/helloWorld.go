package main

import (
	"fmt"
	"log"
	// "rsc.io/quote"
	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including the log entry prefix and a flag to disable printing the time, source file, and line number.
    log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Test the quote module
	// fmt.Println(quote.Glass())
	// fmt.Println(quote.Go())
	// fmt.Println(quote.Hello())
	// fmt.Println(quote.Opt())
	
	fmt.Println(greetings.CustomHello("Matt"))

	// Test error handling
	message, err := greetings.CustomHello("")

	// If an error was returned, print it to the console and exit the program.
    if err != nil {
        log.Fatal(err)
    }

	// If no error was returned, print the returned message to the console.
	fmt.Println(message)
}
