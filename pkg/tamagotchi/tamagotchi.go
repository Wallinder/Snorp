package tamagotchi

import (
	"context"
	"time"
)

const (
	SNORP_DEFAULT = "./sprites/snorp_default.png"
	SNORP_DIRTY   = "./sprites/snorp_dirty.png"
	SNORP_HAPPY   = "./sprites/snorp_happy.png"
	SNORP_HUNGRY  = "./sprites/snorp_hungry.png"
	SNORP_SAD     = "./sprites/snorp_sad.png"
	SNORP_TIRED   = "./sprites/snorp_tired.png"
)

type Tamagotchi struct {
	Name      string
	StartTime time.Time
	LastSleep time.Time
	LastFeed  time.Time
	LastWash  time.Time
	Settings  Settings
	Status    Status
	Send      chan<- Status
}

type Settings struct {
	ReconcileInterval time.Duration
	TimeUntilHungry   time.Duration
	TimeUntilSleepy   time.Duration
	TimeUntilDirty    time.Duration
	TimeUntilDead     time.Duration
}

type Status struct {
	isDead   bool
	Hungry   bool
	Dirty    bool
	Sleeping bool
}

func Start(ctx context.Context, tamagotchi *Tamagotchi) {
	if tamagotchi == nil {
		tamagotchi = newDefaultTamagotchi()
	}
	tamagotchi.startLifeCycle(ctx)
}

func newDefaultTamagotchi() *Tamagotchi {
	return &Tamagotchi{
		Name:      "snorp",
		StartTime: time.Now(),
		LastSleep: time.Now(),
		LastFeed:  time.Now(),
		LastWash:  time.Now(),
		Settings: Settings{
			ReconcileInterval: 5 * time.Minute,
			TimeUntilHungry:   4 * time.Hour,
			TimeUntilSleepy:   6 * time.Hour,
			TimeUntilDirty:    12 * time.Hour,
			TimeUntilDead:     168 * time.Hour,
		},
		Send: make(chan<- Status),
	}
}
