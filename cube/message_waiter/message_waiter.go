package messagewaiter

import (
	"github.com/bwmarrin/discordgo"
	"sync"
	"time"
)

// waiters defines all of the message waiters.
var waiters []*MessageWaiter

// waitersLock is the thread lock for waiters.
var waitersLock = sync.RWMutex{}

// MessageWaiter is the struct used to describe someone waiting for their message.
type MessageWaiter struct {
	ChannelID string
	UserID string
	Channel chan *discordgo.Message
	Done bool
}

// Result is used to cast a message into the channel.
func (w MessageWaiter) Result(msg *discordgo.Message) {
	if w.Done {
		return
	}
	w.Done = true
	w.Channel <- msg
	close(w.Channel)
	waitersLock.Lock()
	for v := range waiters {
		waiters = append(waiters[:v], waiters[v+1:]...)
	}
	waitersLock.Unlock()
}

// WaitForMessage is used to wait for a message. A Timeout of 0 means it will wait forever.
func WaitForMessage(ChannelID string, UserID string, Timeout int) *discordgo.Message {
	Channel := make(chan *discordgo.Message)
	waiter := MessageWaiter{
		ChannelID: ChannelID,
		UserID: UserID,
		Channel: Channel,
	}
	waitersLock.Lock()
	waiters = append(waiters, &waiter)
	waitersLock.Unlock()

	if Timeout != 0 {
		go func() {
			time.Sleep(time.Minute * time.Duration(Timeout))
			if !waiter.Done {
				waiter.Result(nil)
			}
		}()
	}

	j, m := <-Channel
	if m {
		waiter.Done = true
		return j
	}
	return nil
}

// MessageWaitHandler is used to handle incoming messages.
func MessageWaitHandler(msg *discordgo.Message) {
	waitersLock.RLock()
	for _, v := range waiters {
		if v.UserID == msg.Author.ID && v.ChannelID == msg.ChannelID {
			waitersLock.RUnlock()
			v.Result(msg)
			return
		}
	}
	waitersLock.RUnlock()
}
