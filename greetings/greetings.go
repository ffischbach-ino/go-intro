package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func GreetMultiple(names []string) (map[string]string, error) {
	if len(names) < 1 {
		return nil, errors.New("empty name list")
	}

	// initialize a map with the following syntax: make(map[key-type]value-type) -> https://go.dev/blog/maps
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Greet(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Sup, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
