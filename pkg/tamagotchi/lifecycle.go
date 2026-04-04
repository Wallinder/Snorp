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
			newStatus := Status{
				isDead:   t.isDead(),
				Hungry:   t.needFood(),
				Dirty:    t.needWash(),
				Sleeping: t.needSleep(),
			}
			if newStatus != t.Status {
				t.Send <- t.Status
				t.Status = newStatus
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
