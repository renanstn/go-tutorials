package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Setup logs
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get greetings for multiple names
	names := []string{"Renan", "Karine", "Belotti"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

	// Get greetings fom a single name
	message, err := greetings.Hello("Renan")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
