package instamanager

import (
	"sync"
	"time"
)

type TaskStrategy string

const ConstantSearch TaskStrategy = "constant-search"

type Task struct {
	id string

	mu         sync.Locker
	strategy   TaskStrategy
	query      map[string]int64
	timeWindow time.Duration
}

func (t *Task) registerNewTerm(term string) {
	t.mu.Lock()
	t.query[term] = 0
	t.mu.Unlock()
}

func (t *Task) countFind(term string) {
	t.mu.Lock()
	if current, exist := t.query[term]; exist {
		t.query[term] = current + 1
	}
	t.mu.Unlock()
}
