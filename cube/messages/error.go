package messages

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/styles"
)

// Error is used to send a error message.
func Error(Channel *discordgo.Channel, Title string, ErrorMessage string, Session *discordgo.Session) {
	_, _ = Session.ChannelMessageSendComplex(Channel.ID, &discordgo.MessageSend{
		Content: "",
		Embed: &discordgo.MessageEmbed{
			Title:       Title,
			Description: ErrorMessage,
			Color:       styles.Error,
			Author: &discordgo.MessageEmbedAuthor{
				URL:     "https://cubebot.xyz",
				Name:    "Cube",
				IconURL: Session.State.User.AvatarURL(""),
			},
		},
		Tts:   false,
		Files: nil,
		File:  nil,
	})
}
