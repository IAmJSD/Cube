package drops

import "math/rand"

// checkIfDrop is used to decide if to randomly drop.
func checkIfDrop(ChannelID string) bool {
	// Set the baseline chance to 2%.
	Chance := 2

	// Get the messages per minute.
	Count := getMessagesPerMin(ChannelID)
	if Count > 2 {
		// Max out at 2% extra.
		Count = 2
	}
	Chance += Count

	// Get a number between 1 and 100.
	n := rand.Intn(100-1) + 1

	// Return if Chance >= n.
	return Chance >= n
}
