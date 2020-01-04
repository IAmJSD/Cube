package drops

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/styles"
	"math/rand"
	"sync"
)

// dropInfo is the drop information.
type dropInfo struct {
	message *discordgo.Message
	amount  int
	guild   string
}

// DropIDs are all of the ID's which are used in drops.
var dropIDs = map[string]*dropInfo{}

// DropIDsLock is the thread lock for the drops.
var dropIDsLock = sync.RWMutex{}

// randString is used to generate a random string.
func randString(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// ExecuteDrop is used to run a drop.
func ExecuteDrop(ChannelID string, GuildID string, Session *discordgo.Session, Amount int, Prefix string, Description string, ImageURL *string) {
	// Generate the ID.
	ID := randString(5)

	// Generate the embed.
	var Image *discordgo.MessageEmbedImage
	if ImageURL != nil {
		Image = &discordgo.MessageEmbedImage{
			URL: *ImageURL,
		}
	}
	m, err := Session.ChannelMessageSendComplex(ChannelID, &discordgo.MessageSend{
		Content: "",
		Embed: &discordgo.MessageEmbed{
			Description: Description + " To pick up this money, run `" + Prefix + "pick " + ID + "`.",
			Color:       styles.Generic,
			Footer:      nil,
			Image:       Image,
		},
		Tts:   false,
		Files: nil,
		File:  nil,
	})

	// Set the drop ID.
	if err == nil {
		dropIDsLock.Lock()
		dropIDs[ID] = &dropInfo{
			message: m,
			amount:  Amount,
			guild:   GuildID,
		}
		dropIDsLock.Unlock()
	}
}
