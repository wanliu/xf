package xf

import (
	"sync"
)

type EventHandler func(evt *Event)

type Map map[string]interface{}

var mu sync.Mutex
var index int
var fns = make(map[int]func(*Event))

func register(fn func(*Event)) int {
	mu.Lock()
	defer mu.Unlock()
	index++
	for fns[index] != nil {
		index++
	}
	fns[index] = fn
	return index
}

func lookup(i int) func(*Event) {
	mu.Lock()
	defer mu.Unlock()
	return fns[i]
}

func unregister(i int) {
	mu.Lock()
	defer mu.Unlock()
	delete(fns, i)
}
