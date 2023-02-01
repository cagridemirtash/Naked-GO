package main

import (
	g "example.com/greetings"
	f "fmt"
	l "log"
)

func main() {
	//Log package hold some methods
	//SetPrefix is prefix which is throw an error when assign to message
	l.SetPrefix("greetings : ")

	l.SetFlags(0)
	// Get a greeting message and handle error mechanism
	message, err := g.Hello("")
	//If error was found, this if block runs and
	if err != nil {
		l.Fatal(err)
	}
	//If return not a error, print the message on console.
	f.Println(message)
}
