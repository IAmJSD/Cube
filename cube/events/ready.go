package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	guildscount "github.com/jakemakesstuff/Cube/cube/guilds_count"
	"time"
)

func init() {
	// The ready event handler.
	Events = append(Events, func(s *discordgo.Session, event *discordgo.Ready) {
		// Ensure the count is up to date.
		if guildscount.Count() == 0 {
			// This isn't up to date.
			GuildIDS := make([]interface{}, len(s.State.Guilds))
			for i, v := range s.State.Guilds {
				GuildIDS[i] = v.ID
			}
			guildscount.Insert(GuildIDS...)
		}

		// Handle the status.
		go func() {
			for {
				_ = s.UpdateStatus(0, fmt.Sprintf("Cube | Serving %v guilds", guildscount.Count()))
				time.Sleep(time.Second * 10)
			}
		}()
	})
}
