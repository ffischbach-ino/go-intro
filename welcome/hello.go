package main

import (
	"fmt"
	"log"

	"inovex.de/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Fynn")

	if err != nil {
		log.Fatal((err))
	}

	fmt.Println(message)
}
