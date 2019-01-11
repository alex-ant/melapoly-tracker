package params

import (
	"flag"
	"log"

	"github.com/alex-ant/envs"
)

var (
	APIPort = flag.Int("api-port", 30303, "HTTP API port number")
)

func init() {
	// Parse flags if not parsed already.
	if !flag.Parsed() {
		flag.Parse()
	}

	// Determine and read environment variables.
	flagsErr := envs.GetAllFlags()
	if flagsErr != nil {
		log.Fatal(flagsErr)
	}
}
