package tamagotchi

import (
	"context"
	"time"
)

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

func (t *Tamagotchi) startLifeCycle(ctx context.Context) {
	ticker := time.NewTicker(t.Settings.ReconcileInterval)

	var totalControlLoops int64
	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			if t.isDead() {
				t.Dead = true
				return
			}

			if t.needFood() {
			}

			if t.needWash() {
			}

			if t.needSleep() {
			}

			totalControlLoops++
		}
	}
}
