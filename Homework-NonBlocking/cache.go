package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

type cache struct {
	content  map[processID]*procState
	registry map[processID]chan bool
	mu       sync.Mutex
}

func newCache() *cache {
	return &cache{
		content:  make(map[processID]*procState),
		registry: make(map[processID]chan bool),
	}
}

func (c *cache) unregister(id processID) {
	c.mu.Lock()
	delete(c.registry, id)
	c.mu.Unlock()
}

func (c *cache) getregistration(id processID) (chan bool, error) {
	c.mu.Lock()
	res, exists := c.registry[id]
	c.mu.Unlock()

	if !exists {
		return nil, fmt.Errorf("no registration for process ID: %s", id)
	}

	return res, nil
}

func (c *cache) processExists(id processID) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exists := c.content[id]

	return exists
}

func (c *cache) addProcess(p *process, reg chan bool) error {
	if c.processExists(p.id) {
		return fmt.Errorf("add process: process %s exists in cache", p.id)
	}

	c.mu.Lock()
	c.content[p.id] = &p.procState

	c.registry[p.id] = reg
	c.mu.Unlock()

	return nil
}

func (c *cache) updateProcess(p *process) error {
	if !c.processExists(p.id) {
		return fmt.Errorf("update process: process %s does not exist in cache", p.id)
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.content[p.id] = &p.procState

	return nil
}

func (c *cache) getProcState(id processID) *procState {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.content[id]
}

func (c *cache) getProcessState(id processID) (*procState, error) {
	if !c.processExists(id) {
		return nil, fmt.Errorf("get process: process %s does not exist in cache", id)
	}

	p := c.getProcState(id)
	if p.isReady {
		return p, nil
	}

	ch, errReg := c.getregistration(id)
	if errReg != nil {
		return nil, errReg
	}

	<-ch

	return c.getProcState(id), nil
}

func (c cache) getAllProcessIDs() ([]string, error) {
	if len(c.content) == 0 {
		return nil, errors.New("cache is empty")
	}

	var res []string

	for id, proc := range c.content {
		res = append(res, string(id)+" - Is Ready: "+strconv.FormatBool(proc.isReady))
	}

	return res, nil
}

func (c cache) getRegister() int {
	return len(c.registry)
}
