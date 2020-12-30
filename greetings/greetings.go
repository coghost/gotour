package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// implement a HelloBasic function to return greetings
// function whose name starts with a capital letter can be called by a function not in same package
func HelloBasic(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func HelloGuys(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := HelloBasic(name)

		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome",
		"Great to see you, %v",
		"Hail, %v Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
