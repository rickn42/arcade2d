package entity

import . "github.com/rickn42/adventure2d"

type Updater struct {
	f func(this Entity, dt float64)
}

func NewUpdater(f func(this Entity, dt float64)) *Updater {
	return &Updater{f}
}

func (u *Updater) Update(this Entity, dt float64) {
	u.f(this, dt)
}
