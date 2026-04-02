package tamagotchi

import (
	"snorp/pkg/tamagotchi/receiver"
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
	Age       int64
	Awake     bool
	LastSleep time.Time
	LastFeed  time.Time
	OnStart   func()
	OnDirty   func()
	OnTired   func()
	OnHungry  func()
	receiver.Reciever
}

func Start(name string) {
	tamagotchi := newDefaultTamagotchi()
	tamagotchi.Name = name
}

func newDefaultTamagotchi() *Tamagotchi {
	return &Tamagotchi{
		Name:      "gomatchi",
		StartTime: time.Now(),
		Awake:     true,
	}
}
