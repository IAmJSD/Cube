package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/cacher"
)

func init() {
	// The channel update event handler.
	Events = append(Events, func(_ *discordgo.Session, event *discordgo.ChannelUpdate) {
		go cacher.DeleteChannels(event.ID)
	})
}
