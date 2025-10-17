package main

import "fmt"

type Status int
type ServerState int

const (
	Failed Status = iota
	Pending
	Complete
)
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{ //map declared
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (s Status) String() string {
	switch s {
	case Failed:
		return "Failed"
	case Pending:
		return "Pending"
	case Complete:
		return "Complete"
	default:
		return "Unknown"

	}
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:

		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
func Products(productStatus Status) {
	fmt.Printf("The Product status is %v\n", productStatus)

}
func main() {
	Products(Failed)
	Products(Pending)
	Products(Complete)

	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}
