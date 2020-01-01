package drops

import "math/rand"

// checkIfDrop is used to decide if to randomly drop.
func checkIfDrop(ChannelID string) bool {
	// Set the baseline chance to 10%.
	Chance := 10

	// Get the messages per minute.
	count := getMessagesPerMin(ChannelID)
	if count > 15 {
		// Max out at 15% extra.
		count = 15
	}
	Chance += count

	// Get a number between 1 and 100.
	n := rand.Intn(100-1) + 1

	// Return if Chance >= n.
	return Chance >= n
}
