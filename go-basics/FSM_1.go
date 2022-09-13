package main

import "fmt"

// Transitions that can be made from a state
type Transition struct {
	To string
}

// Transition map to store the transition corresponding to the event
type TransitionMap map[string]Transition

// State details for a state
type StateDetails struct {
	On TransitionMap
}

// State and its corresponding transiions depending on event
type StateMap map[string]StateDetails

// State Machine details
type Machine struct {
	ID      string
	Initial string
	current string
	States  StateMap
}

// Returns current state of the machine
func (m *Machine) Current() string {
	if m.current == "" {
		return m.Initial
	}
	return m.current
}

// Make transition from one state to another depending on the event
func (m *Machine) Transition(event string) string {
	current := m.Current()
	transitions := m.States[current].On
	next := transitions[event].To
	if next != "" {
		m.current = next
	}
	return current
}

func main() {
	//////////////////////////////////////////////////////////
	// Initial State : lock
	// Current state		Event			Next State
	// =============		=====			==========
	// lock					coin			unlock
	// lock					push			lock
	// unlock				coin			unlock
	// unlock				push			lock
	//////////////////////////////////////////////////////////

	machine := Machine{
		ID:      "machine-1",
		Initial: "lock",
		States: StateMap{
			"lock": StateDetails{
				On: TransitionMap{
					"coin": Transition{
						To: "unlock",
					},
					"push": Transition{
						To: "lock",
					},
				},
			},

			"unlock": StateDetails{
				On: TransitionMap{
					"coin": Transition{
						To: "unlock",
					},
					"push": Transition{
						To: "lock",
					},
				},
			},
		},
	}

	fmt.Println(machine.Current())
	machine.Transition("coin")
	fmt.Println(machine.Current())
	machine.Transition("push")
	fmt.Println(machine.Current())
}

// https://medium.com/wesionary-team/finite-state-machines-with-go-lang-ccd20e329a7b
