package cube

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/events"
	"github.com/jakemakesstuff/Cube/cube/sharding"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	// Used to initialise Redis.
	_ "github.com/jakemakesstuff/Cube/cube/redis"

	// Used to initialise commands.
	_ "github.com/jakemakesstuff/Cube/cube/commands"
)

// Init is used to initialise the application.
func Init() {
	// Set the random number seed.
	rand.Seed(time.Now().UnixNano())

	// Creates the Discord session.
	session, err := discordgo.New("Bot " + Token)
	if err != nil {
		panic(err)
	}
	session.ShardID = sharding.ShardID
	session.ShardCount = sharding.ShardCount

	// Registers all events.
	events.Register(session)

	// Open the websocket and begin listening.
	err = session.Open()
	if err != nil {
		panic("Error opening Discord session: " + err.Error())
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Cube is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = session.Close()
}
