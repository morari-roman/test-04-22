package scheduler

import (
	"time"

	"github.com/madflojo/tasks"
)

type Service interface {
	GetData() error
}

type Scheduler struct {
	sc          *tasks.Scheduler
	interval    time.Duration
	serviceFunc []Service
}

func New(i time.Duration) *Scheduler {
	return &Scheduler{
		sc:       tasks.New(),
		interval: i,
	}
}

func (sc *Scheduler) AddService(funcTask Service) error {
	sc.serviceFunc = append(sc.serviceFunc, funcTask)

	return nil
}
