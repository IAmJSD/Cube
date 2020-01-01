package reactionwaiter

import (
	"github.com/bwmarrin/discordgo"
	"sync"
	"time"
)

// waiters defines all of the reaction waiters.
var waiters []*ReactionWaiter

// waitersLock is the thread lock for waiters.
var waitersLock = sync.RWMutex{}

// ReactionWaiter is the struct used to describe someone waiting for their reaction.
type ReactionWaiter struct {
	MessageID string
	UserID    string
	Channel   chan *discordgo.Emoji
	Done      bool
}

// Result is used to cast a reaction into the channel.
func (w ReactionWaiter) Result(e *discordgo.Emoji) {
	if w.Done {
		return
	}
	w.Done = true
	w.Channel <- e
	close(w.Channel)
	waitersLock.Lock()
	for v := range waiters {
		waiters = append(waiters[:v], waiters[v+1:]...)
	}
	waitersLock.Unlock()
}

// WaitForReaction is used to wait for a reaction. A Timeout of 0 means it will wait forever.
func WaitForReaction(MessageID string, UserID string, Timeout int) *discordgo.Emoji {
	Channel := make(chan *discordgo.Emoji)
	waiter := ReactionWaiter{
		MessageID: MessageID,
		UserID:    UserID,
		Channel:   Channel,
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

// ReactionWaitHandler is used to handle incoming reactions.
func ReactionWaitHandler(r *discordgo.MessageReactionAdd) {
	waitersLock.RLock()
	for _, v := range waiters {
		if v.UserID == r.UserID && v.MessageID == r.MessageID {
			waitersLock.RUnlock()
			v.Result(&r.Emoji)
			return
		}
	}
	waitersLock.RUnlock()
}
