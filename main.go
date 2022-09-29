package main

import (
	"math/rand"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	rand.Seed(time.Now().UnixNano())

	//	Every 5 seconds ...
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			//	... run our logging code
			log.Info().
				Int("Random number", rand.Intn(1000)).
				Msg("Logging a message")
		case <-quit:
			ticker.Stop()
			return
		}
	}

}
