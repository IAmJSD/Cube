package wallets

import (
	"github.com/jakemakesstuff/Cube/cube/redis"
	"strconv"
)

// SetBalance is used to set a users balance.
func SetBalance(UserID string, GuildID string, Amount int64) {
	Key, Lock := ensureWallet(GuildID, UserID)
	Lock.Lock()
	Formatted := strconv.FormatInt(Amount, 10)
	redis.Client.Set(Key, Formatted, 0)
	Lock.Unlock()
}
