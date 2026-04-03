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
	Dead      bool
	StartTime time.Time
	LastSleep time.Time
	LastFeed  time.Time
	LastWash  time.Time
	Settings  Settings
	Reciever
}

type Settings struct {
	ReconcileInterval time.Duration
	TimeUntilHungry   time.Duration
	TimeUntilSleepy   time.Duration
	TimeUntilDirty    time.Duration
	TimeUntilDead     time.Duration
}

func Start(ctx context.Context, tamagotchi *Tamagotchi) {
	if tamagotchi == nil {
		tamagotchi = newDefaultTamagotchi()
	}
	tamagotchi.startLifeCycle(ctx)
}

func newDefaultTamagotchi() *Tamagotchi {
	return &Tamagotchi{
		Name:      "gomatchi",
		StartTime: time.Now(),
		LastSleep: time.Now(),
		LastFeed:  time.Now(),
		LastWash:  time.Now(),
		Dead:      false,
		Settings: Settings{
			ReconcileInterval: 10 * time.Minute,
			TimeUntilHungry:   4 * time.Hour,
			TimeUntilSleepy:   8 * time.Hour,
			TimeUntilDirty:    4 * time.Hour,
			TimeUntilDead:     168 * time.Hour,
		},
	}
}
