package drops

import (
	"github.com/bwmarrin/discordgo"
	"github.com/getsentry/sentry-go"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/redis"
	"os"
	"strconv"
)

// HandleRandomDrops is used to handle random money drops.
func HandleRandomDrops(msg *discordgo.Message, Session *discordgo.Session) {
	// Get the currency config.
	cur := currency.GetCurrency(msg.GuildID)

	// Check if drops are enabled.
	if !cur.DropsEnabled {
		// They are not! Return here.
		return
	}

	// Check if we should drop.
	ShouldDrop := checkIfDrop(msg.ChannelID)
	if !ShouldDrop {
		// We shouldn't drop - return here!
		return
	}

	// Gets the prefix.
	Prefix := os.Getenv("DEFAULT_PREFIX")
	r, err := redis.Client.Get("p:" + msg.GuildID).Result()
	if err == nil {
		// Set the new prefix.
		Prefix = r
	} else if err != redis.Nil {
		sentry.CaptureException(err)
		panic(err)
	}

	// Run the drop.
	DropsAmount := 100
	if cur.DropsAmount != nil {
		DropsAmount = *cur.DropsAmount
	}
	ExecuteDrop(msg.ChannelID, msg.GuildID, Session, DropsAmount, Prefix, *cur.Emoji+" "+strconv.Itoa(DropsAmount)+" just dropped on the floor.", cur.DropsImage)
}
