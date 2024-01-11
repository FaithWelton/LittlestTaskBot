package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("[GREETINGS]: No Name Provided")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v! I'm Taskbot, it's great to meet you!",
		"Hello %v! I'm Taskbot, it's lovely to see you 👋",
		"Hi There %v! I'm Taskbot!",
		"Howdy %v! Taskbot's the name, Tasks are the game!",
		"Great to see you, %v! I'm Taskbot, how wonderful to have you here 😸",
	}

	return formats[rand.Intn(len(formats))]
}
