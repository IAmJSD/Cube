package sharding

import (
	"os"
	"strconv"
	"strings"
)

// ShardID is the current shard ID.
var ShardID = 0

// ShardCount is the current shard count.
var ShardCount = 1

// Initialises the shard information.
func init() {
	// Deal with the shard count first.
	e := os.Getenv("SHARD_COUNT")
	if e != "" {
		i, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}
		ShardCount = i
	}

	// Now deal with the shard ID (we support POD_NAME and SHARD_ID due to Kubernetes).
	e = os.Getenv("SHARD_ID")
	if e != "" {
		i, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}
		ShardID = i
	}
	e = os.Getenv("POD_NAME")
	if e != "" {
		DashSplit := strings.Split(e, "-")
		Last := DashSplit[len(DashSplit) - 1]
		i, err := strconv.Atoi(Last)
		if err != nil {
			panic(err)
		}
		ShardID = i
	}
}
