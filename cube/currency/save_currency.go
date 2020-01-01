package currency

import (
	"encoding/json"
	"github.com/jakemakesstuff/Cube/cube/redis"
)

// SaveCurrency is used to save the currency.
func SaveCurrency(GuildID string, c *Currency) {
	j, err := json.Marshal(c)
	if err != nil {
		return
	}
	_ = redis.Client.Set("C:"+GuildID, j, 0).Err()
}
