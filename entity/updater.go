package entity

import (
	"time"

	. "github.com/rickn42/adventure2d"
)

type Updater struct {
	f func(this Entity, dt time.Duration)
}

func NewUpdater(f func(this Entity, dt time.Duration)) *Updater {
	return &Updater{f}
}

func (u *Updater) String() string {
	return "Updater"
}

func (u *Updater) Update(this Entity, dt time.Duration) {
	u.f(this, dt)
}
