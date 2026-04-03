package tamagotchi

import "time"

func (t *Tamagotchi) Feed(food string) {
	t.LastFeed = time.Now()
	t.Status.Happy = true
	t.Send <- t.Status
}

func (t *Tamagotchi) Wash(food string) {
	t.LastWash = time.Now()
	t.Status.Dirty = false
	t.Send <- t.Status
}

func (t *Tamagotchi) onHungry() {
	status := Status{}
	t.Send <- status
}

func (t *Tamagotchi) onDirty() {
	status := Status{}
	t.Send <- status
}

func (t *Tamagotchi) onSleep() {
	status := Status{}
	t.Send <- status
}
