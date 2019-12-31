package wallets

import (
	"errors"
	"github.com/jakemakesstuff/Cube/cube/redis"
)

// AddToBalance is used to add to a users balance.
func AddToBalance(UserID string, GuildID string, Amount int64) error {
	Key, Lock := ensureWallet(GuildID, UserID)
	if 0 > Amount {
		return errors.New("The amount you wish to add cannot be negative. ")
	}
	Lock.Lock()
	redis.Client.IncrBy(Key, Amount)
	Lock.Unlock()
	return nil
}
