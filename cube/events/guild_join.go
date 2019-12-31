package events

import (
	"github.com/bwmarrin/discordgo"
	guildscount "github.com/jakemakesstuff/Cube/cube/guilds_count"
)

func init() {
	// The guild create event handler.
	Events = append(Events, func(_ *discordgo.Session, event *discordgo.GuildCreate) {
		// Add this guild to the count.
		go guildscount.Insert(event.ID)
	})
}
