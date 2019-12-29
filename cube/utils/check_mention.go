package utils

import "github.com/bwmarrin/discordgo"

// CheckMention is used to check if the bot was mentioned.
func CheckMention(Content string, session *discordgo.Session) int {
	// Get the bot ID.
	BotID := session.State.User.ID

	// Check the length.
	if len(BotID) > len(Content) {
		// The length of the bot ID is greater than the content.
		// It can't be a mention!
		return 0
	}

	// Check the chars.
	BeginningChars := ""
	ID := ""
	ExtraStuffLen := 2
	for i, v := range Content {
		if 2 > i {
			// Add to the beginning chars.
			BeginningChars += string(v)
		} else if i == 2 {
			if BeginningChars != "<@" {
				// Invalid ID.
				return 0
			} else {
				// Check if this is a special char or part of the ID.
				switch v {
				case '!':
					// Nickname.
					ExtraStuffLen++
					break
				case '&':
					// Role.
					return 0
				default:
					// Add to the ID.
					ID += string(v)
					break
				}
			}
		} else {
			// Add the remainder of the ID.
			if v == '>' {
				// The end of the ID.
				ExtraStuffLen++
				break
			} else {
				// Add to the rest of the ID.
				ID += string(v)
			}
		}
	}

	// Return the length.
	return len(ID) + ExtraStuffLen
}
