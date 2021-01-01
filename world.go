package main

import (
	"log"
	"time"
)

// World contains the model of the game world
type World struct {
	worldName         string
	worldTime         int
	worldCreationTime time.Time
	worldPopulation   int
	worldType         string
	worldTicker       *time.Ticker
	playerList        []string
}

func newWorld(name string, worldType string) *World {
	log.Printf("Creating the world %s", name)
	return &World{
		worldName:         name,
		worldTime:         0,
		worldCreationTime: time.Now(),
		worldPopulation:   0,
		worldType:         worldType,
		worldTicker:       time.NewTicker(1 * time.Second),
	}

}
func joinWorld() {

}
func (world *World) run() {
	for {
		select {
		case <-world.worldTicker.C:
			world.worldTime++
			log.Printf("World Time passed in sec %d", world.worldTime)
		}
	}
}
