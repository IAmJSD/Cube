package roleshop

import (
	"sync"
	"time"
)

// trialMaps are the map of trials.
var trialMaps = map[string]*map[string]*chan bool{}

// trialMapsLock is the thread lock for trialMaps.
var trialMapsLock = sync.Mutex{}

// waitForTrial is used to wait until a trial is over. The boolean represents if to delete the role from the user or not.
func waitForTrial(RoleID string, UserID string) bool {
	// Lock the trials map lock.
	trialMapsLock.Lock()

	// Get the users map.
	Users, ok := trialMaps[RoleID]
	if !ok {
		// Create the users map.
		m := map[string]*chan bool{}
		Users = &m
		trialMaps[RoleID] = Users
	}

	// Create the channel if it does not exist.
	c, ok := (*Users)[UserID]
	var channel chan bool
	if ok {
		channel = *c
	} else {
		channel = make(chan bool)
		(*Users)[UserID] = &channel

		// Create a thread to handle the countdown.
		go func() {
			// Sleep for a min.
			time.Sleep(time.Minute)

			// nil-ify the channel if it's the same one.
			trialMapsLock.Lock()
			if (*Users)[UserID] == &channel {
				// Remove the channel from the map.
				delete(*Users, UserID)
			}
			trialMapsLock.Unlock()

			// Return true on the channel.
			channel <- true
		}()
	}

	// Unlock the trials map lock.
	trialMapsLock.Unlock()

	// Wait for the trial result.
	return <-channel
}

// handleBuyTrial is used to handle a person buying the role with the trial controller.
func handleBuyTrial(RoleID string, UserID string) {
	// Lock the trials map lock.
	trialMapsLock.Lock()

	// Get the users map.
	Users, ok := trialMaps[RoleID]
	if !ok {
		// Create the users map.
		m := map[string]*chan bool{}
		Users = &m
		trialMaps[RoleID] = Users
	}

	// Check if it is a user channel.
	Channel, ok := (*Users)[UserID]
	if ok {
		// Remove the channel from the map.
		delete(*Users, UserID)

		// Return false on the channel.
		*Channel <- false
	}

	// Unlock the trials map lock.
	trialMapsLock.Unlock()
}
