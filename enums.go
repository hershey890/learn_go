package main

import "fmt"

type ServerState int

const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func transition(ss ServerState) ServerState {
	switch ss {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", ss))
	}

	return StateConnected
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	ns2 := transition(StateConnected)
	fmt.Println(ns2)
}
