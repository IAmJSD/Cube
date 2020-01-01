package drops

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/wallets"
	"strconv"
	"time"
)

// PickUp is used to pick up the currency.
func PickUp(ID string, ChannelID string, MessageID string, GuildID string, UserID string, UserMention string, Currency *currency.Currency, Session *discordgo.Session) {
	// Delete the message if we can.
	_ = Session.ChannelMessageDelete(ChannelID, MessageID)

	// Get the drop.
	dropIDsLock.RLock()
	drop := dropIDs[ID]
	dropIDsLock.RUnlock()

	// If drop is nil, return here.
	if drop == nil {
		return
	}

	// Delete the drop.
	dropIDsLock.Lock()
	delete(dropIDs, ID)
	dropIDsLock.Unlock()

	// Add the amount to the users balance.
	_ = wallets.AddToBalance(UserID, GuildID, int64(drop.amount))

	// Delete the drop message if possible.
	_ = Session.ChannelMessageDelete(drop.message.ChannelID, drop.message.ID)

	// Send a message saying that the user picked up money.
	msg, err := Session.ChannelMessageSendComplex(ChannelID, &discordgo.MessageSend{
		Content: UserMention + " picked up " + strconv.Itoa(drop.amount) + " " + *Currency.Emoji + "!",
	})
	if err == nil {
		go func() {
			time.Sleep(time.Second * 5)
			_ = Session.ChannelMessageDelete(msg.ChannelID, msg.ID)
		}()
	}
}
