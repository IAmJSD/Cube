package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/embed_menus"
	"github.com/jakemakesstuff/Cube/cube/reaction_waiter"
)

func init() {
	// Defines the reaction add handler.
	Events = append(Events, func(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
		// Handle reaction waiters.
		go reactionwaiter.ReactionWaitHandler(m)

		go func() {
			// Handle embed menus.
			message, err := s.ChannelMessage(m.ChannelID, m.MessageID)
			if err != nil {
				return
			}
			user, err := s.User(m.UserID)
			if err != nil {
				return
			}
			if user.Bot {
				return
			}
			if message.Author.ID == s.State.User.ID && len(message.Embeds) == 1 && message.Embeds[0].Footer != nil {
				MenuID := message.Embeds[0].Footer.Text
				embedmenus.HandleMenuReactionEdit(s, m, MenuID)
			}
		}()
	})
}
