package instamanager

import (
	"sync"
	"time"
)

type TaskKind string

const PendingTask TaskKind = "pending"
const ProcessingTask TaskKind = "processing"

// const ATask TaskKind = "processing"

type Cluster struct {
	mu sync.Locker
	// workers map[*Worker]WorkerState

	tasks      map[TaskKind][]*Task
	assignees  map[string]*Worker
	startTimes map[string]time.Time
}

func newCluster() *Cluster {
	return &Cluster{
		mu: &sync.Mutex{},
		// workers:   map[*Worker]WorkerState{},
		assignees:  map[string]*Worker{},
		startTimes: map[string]time.Time{},
		tasks: map[TaskKind][]*Task{
			PendingTask:    {},
			ProcessingTask: {},
		},
	}
}

func (c *Cluster) registerNewTask(task *Task) {
	c.mu.Lock()
	c.tasks[PendingTask] = append(c.tasks[PendingTask], task)
	c.mu.Unlock()

	// go c.taskLoop(task, 0)
}

func (c *Cluster) requestTask(w *Worker) {
	for i, task := range c.tasks[PendingTask] {
		c.mu.Lock()
		c.assignees[task.id] = w
		// c.workers[w] = WorkerBusy

		c.tasks[ProcessingTask] = append(c.tasks[ProcessingTask], task)

		c.tasks[PendingTask] = append(
			c.tasks[PendingTask][:i],
			c.tasks[PendingTask][i+1:]...,
		)

		c.startTimes[task.id] = time.Now()
		c.mu.Unlock()
		break
	}
}

func (c *Cluster) resolveTask(w *Worker) {
	
}

// func (c *Cluster) taskLoop(task *Task, times int) {
// 	startTime := time.Now()

// 	for w, state := range c.workers {
// 		c.mu.Lock()
// 		if state == WorkerIdle {
// 			c.workers[w] = WorkerBusy
// 			c.assignees[task.id] = w
// 		}
// 		c.mu.Unlock()

// 		go c.xyz(startTime, task)

// 		return
// 	}

// 	time.Sleep(1 * time.Second)

// 	c.taskLoop(task, times+1)
// }

// func (c *Cluster) xyz(startTime time.Time, task *Task) {

// }
