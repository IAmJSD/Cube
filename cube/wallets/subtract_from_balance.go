package wallets

import (
	"errors"
	"github.com/jakemakesstuff/Cube/cube/redis"
)

// SubtractFromBalance is used to subtract from a users balance.
func SubtractFromBalance(UserID string, GuildID string, Amount int64) error {
	Key, Lock := ensureWallet(GuildID, UserID)
	if 0 > Amount {
		return errors.New("The amount you wish to subtract cannot be negative. ")
	}
	Lock.Lock()
	redis.Client.IncrBy(Key, Amount*-1)
	Lock.Unlock()
	return nil
}
