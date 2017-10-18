package system

import (
	"fmt"
	"io"
	"time"

	. "github.com/rickn42/arcade2d"
)

type watcher struct {
	order   int
	watchDt time.Duration
	dt      time.Duration
	w       io.Writer
	c       chan string
}

func WatcherSystem(writer io.Writer, watchDt time.Duration) *watcher {
	c := make(chan string, 100)
	go func() {
		for s := range c {
			writer.Write([]byte(s))
		}
	}()
	return &watcher{
		watchDt: watchDt,
		w:       writer,
		c:       c,
	}
}

func (w *watcher) Order() int {
	return w.order
}

func (w *watcher) SetOrder(n int) *watcher {
	w.order = n
	return w
}

func (w *watcher) Add(Entity) error { return nil }

func (w *watcher) Remove(Entity) {}

func (w *watcher) Update(es []Entity, dt time.Duration) {
	if w.watchDt != 0 {
		w.dt += dt
		if w.dt < w.watchDt {
			return
		}
		w.dt = 0
	}
	w.Watch(es)
}

func (w *watcher) Watch(es []Entity) {
	w.c <- fmt.Sprintf("Watch entities (count=%d)\n", len(es))

	for _, e := range es {
		w.c <- fmt.Sprintf("%T %v\n", e, e)
	}

	w.c <- "\n"
}
