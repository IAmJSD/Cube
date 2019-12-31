package messages

import (
	"github.com/bwmarrin/discordgo"
)

// NotEnoughArgs is the "Not enough arguments" message which the bot sends.
func NotEnoughArgs(Channel *discordgo.Channel, Session *discordgo.Session) {
	Error(Channel, "Not enough arguments:", "There were not enough arguments to use this command.", Session)
}
