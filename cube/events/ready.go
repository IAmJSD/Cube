package events

import "github.com/bwmarrin/discordgo"

func init() {
	// The ready event handler.
	Events = append(Events, func(s *discordgo.Session, event *discordgo.Ready) {
		_ = s.UpdateStatus(0, "Cube | c!help for help with the bot")
	})
}
