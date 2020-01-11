package drops

import "math/rand"

// checkIfDrop is used to decide if to randomly drop.
func checkIfDrop(ChannelID string) bool {
	// Set the baseline chance to 4%.
	Chance := 4

	// Get the messages per minute.
	count := getMessagesPerMin(ChannelID)
	if count > 5 {
		// Max out at 5% extra.
		count = 5
	}
	Chance += count

	// Get a number between 1 and 100.
	n := rand.Intn(100-1) + 1

	// Return if Chance >= n.
	return Chance >= n
}
