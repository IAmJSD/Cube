package messages

import (
	"github.com/bwmarrin/discordgo"
)

// NotAPositiveInteger is the "Not a positive integer" message which the bot sends.
func NotAPositiveInteger(Channel *discordgo.Channel, Session *discordgo.Session, Arg string) {
	Error(Channel, "Not a integer:", "The argument \""+Arg+"\" is meant to be a positive integer.", Session)
}
