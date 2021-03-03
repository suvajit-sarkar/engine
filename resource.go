package main

import (
	"log"
	"time"
)

// ResourceEngine is the engine for the game world
type ResourceEngine struct {
	gold   int
	ticker *time.Ticker
}

func newResourceEngine() *ResourceEngine {
	return &ResourceEngine{
		ticker: time.NewTicker(1 * time.Second),
		gold:   0,
	}
}

func (re *ResourceEngine) run() {
	for {
		select {
		case <-re.ticker.C:
			re.gold++
			log.Println(re.gold)
		}
	}
}
