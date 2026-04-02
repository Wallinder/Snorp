package tamagotchi

import (
	"snorp/pkg/tamagotchi/receiver"
	"time"
)

type Tamagotchi struct {
	Name      string
	StartTime time.Time
	Age       int64
	Awake     bool
	LastSleep time.Time
	LastFeed  time.Time
	OnStart   func()
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
