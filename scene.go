package adventure2d

import (
	"errors"
	"reflect"
	"sort"
	"time"

	"github.com/murlokswarm/log"
)

type Scene struct {
	id        EntityID
	frameRate int
	ss        systems
	es        []Entity
	genID     func() EntityID
}

func (e *Scene) ID() EntityID {
	return e.id
}

func (e *Scene) SetID(id EntityID) {
	e.id = id
}

func (s *Scene) AddEntity(e Entity) error {
	err := s.nilPropertyCheck(e)
	if err != nil {
		log.Error("ADD ENTITY FAILED\n", err)
		return err
	}

	e.SetID(s.genID())
	s.es = append(s.es, e)
	log.Infof("ADD ENTITY\n%v\n", e)

	for _, system := range s.ss {
		err := system.Add(e)
		if err != nil {
			EventManager.Send(err)
		}
	}

	return nil
}

var EntityHasNilProperties = errors.New("Entity has <nil> properties")

func (s *Scene) nilPropertyCheck(e Entity) error {
	rv := reflect.ValueOf(e)
	switch rv.Kind() {
	case reflect.Ptr:
		rv = rv.Elem()
	}

	for i := 0; i < rv.NumField(); i++ {
		rf := rv.Field(i)
		switch rf.Kind() {
		case reflect.Interface, reflect.Ptr:
			if rv.Field(i).IsNil() {
				return EntityHasNilProperties
			}
		}
	}
	return nil
}

func (s *Scene) AddSystem(ss ...System) {
	s.ss = append(s.ss, ss...)
	sort.Sort(s.ss)
}

func (s *Scene) Play() {

	dt := UnitDt / time.Duration(s.frameRate)

	for {
		for _, system := range s.ss {
			system.Update(s.es, dt)
		}
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
