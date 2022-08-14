package main

import (
	"fmt"
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
)

type cfgProc struct {
	id          processID
	c           *cache
	etaMilisecs int
}

type procState struct {
	output  string
	isReady bool
}

type processID string

type process struct {
	procState

	id processID
}

var generatorProcID = func() processID { return processID(uuid.NewV4().String()) }

func newProcess(cfg cfgProc) *process {
	p := process{
		id: cfg.id,
	}

	chReady := make(chan bool)

	go cfg.c.addProcess(&p, chReady)

	p.doWork(cfg.etaMilisecs)

	go func() {
		cfg.c.updateProcess(&p)

		go cfg.c.unregister(p.id)

		ticker := time.NewTicker(300 * time.Millisecond)

		select {
		case chReady <- true:
			{
				close(chReady)
			}

		case <-ticker.C:
			close(chReady)
		}
	}()

	return &p
}

func (p *process) doWork(etaMilisecs int) error {
	rand.Seed(time.Now().UnixNano())

	workDuration := rand.Intn(etaMilisecs)
	time.Sleep(time.Duration(workDuration) * time.Millisecond)

	p.procState = procState{
		isReady: true,
		output:  fmt.Sprintf("process %s took %d miliseconds.", p.id, workDuration),
	}

	return nil
}
