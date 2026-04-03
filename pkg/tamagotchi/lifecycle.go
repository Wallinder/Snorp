package tamagotchi

import (
	"context"
	"time"
)

func (t *Tamagotchi) startLifeCycle(ctx context.Context) {
	ticker := time.NewTicker(t.Settings.ReconcileInterval)

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			t.Status.isDead = t.isDead()
			if t.needFood() {
				t.onHungry()
			}
			if t.needWash() {
				t.onDirty()
			}
			if t.needSleep() {
				t.onSleep()
			}
		}
	}
}

func (t *Tamagotchi) needSleep() bool {
	return time.Since(t.LastSleep) > t.Settings.TimeUntilSleepy
}

func (t *Tamagotchi) needWash() bool {
	return time.Since(t.LastWash) > t.Settings.TimeUntilDirty
}

func (t *Tamagotchi) needFood() bool {
	return time.Since(t.LastFeed) > t.Settings.TimeUntilHungry
}

func (t *Tamagotchi) isDead() bool {
	return time.Since(t.LastFeed) > t.Settings.TimeUntilDead
}
