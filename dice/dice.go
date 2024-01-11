package dice

import (
	"fmt"
	"math/rand"
)

func Roll(dice int, sides int) (string, error) {
	rolls := "Results:\n"
	for i := 1; i <= dice; i++ {
		min := 1
		max := sides

		// Add 1 to rolls to account for 0 index
		roll := rand.Intn((max-min)+min) + 1
		rolls += fmt.Sprintf("%d: %d\n", i, roll)
	}

	message := fmt.Sprintf("Rolling %dd%d!\n%s", dice, sides, rolls)

	return message, nil
}
