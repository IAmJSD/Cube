package utils

// Mention is the user mention specified.
type Mention struct {
	Len int
	ID  string
}

// CheckMention is used to check if the bot was mentioned.
func CheckMention(Content string, OtherID *string) *Mention {
	// Check the length.
	if OtherID != nil {
		if len(*OtherID) > len(Content) {
			// The length of the bot ID is greater than the content.
			// It can't be a mention!
			return &Mention{}
		}
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
				return &Mention{}
			} else {
				// Check if this is a special char or part of the ID.
				switch v {
				case '!':
					// Nickname.
					ExtraStuffLen++
					break
				case '&':
					// Role.
					return &Mention{}
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

	// Check the ID (if specified).
	if OtherID != nil {
		if ID != *OtherID {
			return &Mention{}
		}
	}

	// Return the length.
	return &Mention{
		Len: len(ID) + ExtraStuffLen,
		ID:  ID,
	}
}
