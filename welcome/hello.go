package main

import (
	"fmt"
	"log"

	"inovex.de/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Hans", "Greg", "Leon"}
	messages, err := greetings.GreetMultiple(names)

	if err != nil {
		log.Fatal((err))
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
