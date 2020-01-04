package events

import (
	"github.com/bwmarrin/discordgo"
	guildscount "github.com/jakemakesstuff/Cube/cube/guilds_count"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"github.com/jakemakesstuff/Cube/cube/redis"
	"github.com/jakemakesstuff/Cube/cube/utils"
)

func init() {
	// The guild create event handler.
	Events = append(Events, func(session *discordgo.Session, event *discordgo.GuildCreate) {
		// Check if this is a new guild.
		if !redis.Client.SIsMember("guilds", event.ID).Val() {
			// Add this guild to the count.
			go guildscount.Insert(event.ID)

			// Send a introduction message.
			go func() {
				channel := utils.FindDefaultChannel(event.Guild)
				if channel != "" {
					for _, v := range event.Guild.Channels {
						if v.ID == channel {
							messages.GenericText(v, "Welcome to Cube!",
								"Welcome to Cube, I'm sure you will love the bot as much as we loved making it."+
									" To get a list of commands you can use with this bot, use `c!help`."+
									"\n\nTo get help with the bot, do not hesitate to join our [support Discord](https://discord.gg/CxPxNg9) or make an issue on the "+
									"[GitHub page](https://github.com/JakeMakesStuff/Cube).\n\n-Cube Development Team", nil, session)
						}
					}
				}
			}()
		}
	})
}
