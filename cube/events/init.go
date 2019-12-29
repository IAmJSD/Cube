package events

import "github.com/bwmarrin/discordgo"

// Events are all of the Discord events.
var Events = make([]interface{}, 0)

// Register is used to register the events into the session.
func Register(session *discordgo.Session) {
	for _, v := range Events {
		session.AddHandler(v)
	}
}
