package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/cacher"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/drops"
	"github.com/jakemakesstuff/Cube/cube/message_waiter"
	"time"
)

func init() {
	// Defines the message handler.
	Events = append(Events, func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Check if the author is a bot and return if so.
		if m.Author.Bot {
			return
		}

		// Handle message waiters.
		go messagewaiter.MessageWaitHandler(m.Message)

		// Handle random drops.
		go drops.HandleRandomDrops(m.Message, s)

		go func() {
			// Gets the start time.
			StartTime := time.Now()

			// Gets the channel.
			channel, err := cacher.GetChannel(m.ChannelID, s)
			if err != nil {
				// TODO: Error report here to Sentry.
				return
			}

			// If this is a DM, ignore it.
			if channel.Type == discordgo.ChannelTypeGroupDM || channel.Type == discordgo.ChannelTypeDM {
				return
			}

			// Fire up the commands processor in a new thread.
			go commandprocessor.Processor(m.Message, channel, s, &StartTime)
		}()
	})
}
