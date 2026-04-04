package tamagotchi

import "time"

func (t *Tamagotchi) Feed(food string) {
	t.LastFeed = time.Now()
	t.Send <- t.Status
}

func (t *Tamagotchi) Wash(food string) {
	t.LastWash = time.Now()
	t.Status.Dirty = false
	t.Send <- t.Status
}
