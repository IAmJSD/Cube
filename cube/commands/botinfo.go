package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/guilds_count"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"github.com/jakemakesstuff/Cube/cube/redis"
	"github.com/jakemakesstuff/Cube/cube/utils"
	"runtime"
)

func init() {
	commandprocessor.Commands["botinfo"] = &commandprocessor.Command{
		Description: "Shows information about the bot.",
		Category:    categories.INFORMATIONAL,
		Function: func(Args *commandprocessor.CommandArgs) {
			// Get runtime info.
			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			// Render the information.
			messages.GenericText(
				Args.Channel, "Cube Information:",
				"Cube is a open source bot hosted by JakeMakesStuff#0001.",
				[]*discordgo.MessageEmbedField{
					{
						Name:   "Library Version:",
						Value:  "discordgo " + discordgo.VERSION,
						Inline: true,
					},
					{
						Name:   "Bot Version:",
						Value:  "Cube " + utils.Version,
						Inline: true,
					},
					{
						Name:   "GitHub Repository:",
						Value:  "[JakeMakesStuff/Cube](https://github.com/JakeMakesStuff/Cube)",
						Inline: true,
					},
					{
						Name:   "Website:",
						Value:  "[https://cubebot.xyz](https://cubebot.xyz)",
						Inline: true,
					},
					{
						Name:   "Guilds:",
						Value:  fmt.Sprintf("%v", guildscount.Count()),
						Inline: true,
					},
					{
						Name:   "Memory Allocated:",
						Value:  fmt.Sprintf("%v MB", m.Alloc/1024/1024),
						Inline: true,
					},
					{
						Name:   "Redis Keys:",
						Value:  fmt.Sprintf("%v", redis.Client.DbSize().Val()),
						Inline: true,
					},
				}, Args.Session,
			)
		},
	}
}
