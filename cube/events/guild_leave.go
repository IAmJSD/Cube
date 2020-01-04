package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/cacher"
	guildscount "github.com/jakemakesstuff/Cube/cube/guilds_count"
	"github.com/jakemakesstuff/Cube/cube/redis"
)

func init() {
	// The guild leave event handler.
	Events = append(Events, func(_ *discordgo.Session, event *discordgo.GuildDelete) {
		if !event.Unavailable {
			// Delete the channels from the cache.
			Channels := make([]string, len(event.Channels))
			for i, v := range event.Guild.Channels {
				Channels[i] = v.ID
			}
			go cacher.DeleteChannels(Channels...)

			// Removes the guild from the count.
			go guildscount.Remove(event.ID)

			// Delete the prefix.
			go redis.Client.Del("p:" + event.ID)
		}
	})
}
