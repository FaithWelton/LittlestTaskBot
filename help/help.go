package help

import "fmt"

func New() (string, error) {
	tasks := map[string]string{
		"hello":    "👋 Bot responds with a random greeting",
		"weather":  "⛅️ Bot asks for location to send appropriate weather data",
		"roll":     "🎲 Bot will roll a die based on data received.\nExpects the format #d#, where the first # is the number of die, and second # is number of sides for the die",
		"help":     "🆘 Bot responds with the list of available tasks and a description - You are here!",
		"settings": "⚙️ Shows Bot settings and commands to edit them",
		"test":     "This is a test",
	}

	helpMsg := "The commands that I know how to interact with are as follows:"
	for k, v := range tasks {
		helpMsg = fmt.Sprintf("%s\n\n/%s: %s\n", helpMsg, k, v)
	}

	message := helpMsg
	return message, nil
}
