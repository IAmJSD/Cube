package drops

import (
	"sync"
	"time"
)

// messagesPerMin is a map of channels > messages per min (roughly!)
var messagesPerMin = map[string]int{}

// messagesPerMinLock is the thread lock for this map.
var messagesPerMinLock = sync.Mutex{}

// getMessagesPerMin is used to add one to the messages per min and return it.
func getMessagesPerMin(ChannelID string) int {
	messagesPerMinLock.Lock()
	Messages := messagesPerMin[ChannelID]
	Messages++
	messagesPerMin[ChannelID] = Messages
	messagesPerMinLock.Unlock()
	go func() {
		time.Sleep(time.Minute)
		messagesPerMinLock.Lock()
		messagesPerMin[ChannelID]--
		messagesPerMinLock.Unlock()
	}()
	return Messages
}
