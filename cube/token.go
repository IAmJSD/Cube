package cube

import "os"

// Token defines the Discord token.
var Token string

func init() {
	Token = os.Getenv("TOKEN")
	if Token == "" {
		panic("TOKEN is nil.")
	}
}
