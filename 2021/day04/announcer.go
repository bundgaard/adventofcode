package main

import (
	"sync"

	"github.com/google/uuid"
)

type Announcer struct {
	mu      sync.Mutex
	clients map[string]chan int
}

func (a *Announcer) Subscribe() (string, chan int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	id := uuid.NewString()
	a.clients[id] = make(chan int)
	return id, a.clients[id]
}

func (a *Announcer) Unsubscribe(id string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	delete(a.clients, id)
}

func (a *Announcer) Broadcast(n int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, ch := range a.clients {
		ch <- n
	}
}
