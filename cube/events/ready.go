package events

import "github.com/bwmarrin/discordgo"

func init() {
	// The ready event handler.
	Events = append(Events, func(s *discordgo.Session, event *discordgo.Ready) {
		_ = s.UpdateStatusComplex(discordgo.UpdateStatusData{
			IdleSince: nil,
			Game:      nil,
			AFK:       false,
			Status:    "Cube | c!help for help with the bot",
		})
	})
}
