package messages

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/styles"
)

// GenericText is used to send a generic message with text in.
func GenericText(Channel *discordgo.Channel, Title string, Description string, Fields []*discordgo.MessageEmbedField, Session *discordgo.Session) {
	_, _ = Session.ChannelMessageSendComplex(Channel.ID, &discordgo.MessageSend{
		Content: "",
		Embed:   &discordgo.MessageEmbed{
			Title:       Title,
			Description: Description,
			Color:       styles.Generic,
			Author:      &discordgo.MessageEmbedAuthor{
				URL:          "https://cubebot.xyz",
				Name:         "Cube",
				IconURL:      Session.State.User.AvatarURL(""),
			},
			Fields: Fields,
		},
		Tts:     false,
		Files:   nil,
		File:    nil,
	})
}
