package utils

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/messages"
)

// ProcessMemberArg is used to process a member argument and throw an error if needed.
func ProcessMemberArg(Channel *discordgo.Channel, Session *discordgo.Session, Arg string) *discordgo.Member {
	// Will be ran if a user was not found.
	UserNotFound := func() {
		messages.Error(Channel, "User not found:", "The user you specified was not found.", Session)
	}

	// Check for a mention.
	Mention := CheckMention(Arg, nil)

	// Check if there is a mention or not.
	if Mention.Len == 0 {
		// There's not.
		UserNotFound()
		return nil
	}

	// Now find the user.
	u, err := Session.GuildMember(Channel.GuildID, Mention.ID)
	if err != nil {
		UserNotFound()
		return nil
	}
	return u
}
