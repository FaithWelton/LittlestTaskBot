package help

import "fmt"

func New() (string, error) {
	tasks := map[string]string{
		"hello":   "Bot will respond with a random greeting",
		"weather": "Bot will request location data to send weather data",
		"roll":    "Bot will roll die as specified in the format #d# \nWith the first number being the number of die, and the second number being the number of sides.",
		"test":    "Command for testing things that don't have a name yet",
		"help":    "You are here - This will display a list of commands",
	}

	helpMsg := "The commands that I know how to interact with are as follows:"
	for k, v := range tasks {
		helpMsg = fmt.Sprintf("%s\n%s - %s", helpMsg, k, v)
	}

	message := helpMsg
	return message, nil
}
