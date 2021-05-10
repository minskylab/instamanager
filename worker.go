package instamanager

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type WorkerState string

const WorkerIdle WorkerState = "idle"
const WorkerBusy WorkerState = "busy"

type Worker struct {
	id string

	name  string
	agent string
	ip    string

	registeredTime time.Time

	computationTime time.Duration
	// state WorkerState
}

func newWorker(name, agent, ip string) *Worker {
	return &Worker{
		id:             uuid.NewV4().String(),
		name:           name,
		agent:          agent,
		ip:             ip,
		registeredTime: time.Now(),
		// state:          WorkerIdle,
	}
}

// func (w *Worker)

// func (w *Worker)
