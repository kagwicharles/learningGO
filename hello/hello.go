package main

import (
	"fmt"
	"log"

	"com.kagwi/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0) 

	names := []string{
		"Kagwi",
		"Gerald",
		"Nora",
	}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}