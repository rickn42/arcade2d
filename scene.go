package adventure2d

import (
	"sort"
	"time"
)

type Scene struct {
	Entity
	frameRate    int
	renderSystem System
	systems
	entities []Entity
	genID    func() EntityID
}

func (s *Scene) AddEntity(e Entity) {
	e.setID(s.genID())
	s.entities = append(s.entities, e)

	for _, system := range s.systems {
		err := system.Add(e)
		if err != nil {
			EventManager.Send(err)
		}
	}
}

func (s *Scene) AddSystem(ss ...System) {
	s.systems = append(s.systems, ss...)
	sort.Sort(s.systems)
}

func (s *Scene) Play() {
	dt := UnitDt / time.Duration(s.frameRate)
	ratioDt := float64(dt) / float64(UnitDt)

	for {
		// update
		for _, system := range s.systems {
			system.Update(s.entities, ratioDt)
		}

		// render
		s.renderSystem.Update(s.entities, ratioDt)

		time.Sleep(dt)
	}
}

// sortable interface
type systems []System

func (ss systems) Len() int {
	return len(ss)
}

func (ss systems) Less(i, j int) bool {
	return ss[i].Order() <= ss[j].Order()
}

func (ss systems) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
