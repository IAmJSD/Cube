package wallets

import (
	"github.com/jakemakesstuff/Cube/cube/redis"
	"strconv"
)

// GetBalance is used to get a users balance.
func GetBalance(UserID string, GuildID string) int {
	Key, Lock := ensureWallet(GuildID, UserID)
	Lock.Lock()
	v := redis.Client.Get(Key).Val()
	Lock.Unlock()
	i, _ := strconv.Atoi(v)
	return i
}
