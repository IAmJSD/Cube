package messages

import (
	"github.com/bwmarrin/discordgo"
)

// NotAnInteger is the "Not an integer" message which the bot sends.
func NotAnInteger(Channel *discordgo.Channel, Session *discordgo.Session, Arg string) {
	Error(Channel, "Not an integer:", "The argument \"" + Arg + "\" is meant to be an integer.", Session)
}
