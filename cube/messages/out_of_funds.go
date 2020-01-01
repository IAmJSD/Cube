package messages

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
)

// OutOfFunds is the "Out of funds" message which the bot sends.
func OutOfFunds(Channel *discordgo.Channel, Session *discordgo.Session, Emoji string, Amount int) {
	Error(Channel, "Out of funds:", "You currently have "+strconv.Itoa(Amount)+Emoji+".", Session)
}
