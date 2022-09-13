package main

import "fmt"

type Transition struct {
	To string
}
type TransitionMap map[string]Transition

type StateDetails struct {
	On TransitionMap
}
type StateMap map[string]StateDetails

type Machine struct {
	Id      string
	Initial string
	current string
	States  StateMap
}

func (m *Machine) CurrentState() string {
	if m.current == "" {
		return m.Initial
	}
	return m.current
}

func (m *Machine) TransitionTo(event string) string {
	current := m.current
	if current == "" {
		return m.Initial
	}
	next := m.States[current].On[event].To
	return next
}

func main() {
	machine := &Machine{
		Id:      "Machine-1",
		Initial: "noOrder",
		States: StateMap{
			"noOrder": StateDetails{
				On: TransitionMap{
					"order": Transition{To: "orderCreating"},
				},
			},
			"orderCreating": StateDetails{
				On: TransitionMap{
					"fail":  Transition{To: "orderFailed"},
					"place": Transition{To: "orderPlaced"},
				},
			},
			"orderFailed": StateDetails{
				On: TransitionMap{
					"create": Transition{To: "orderCreating"},
				},
			},
			"orderPlaced": StateDetails{
				On: TransitionMap{
					"charge": Transition{To: "chargingCart"},
				},
			},
			"chargingCart": StateDetails{
				On: TransitionMap{
					"fail":   Transition{To: "transactionFailed"},
					"charge": Transition{To: "chargingCart"},
					"ship":   Transition{To: "orderShipped"},
				},
			},
		},
	}
	fmt.Println(machine.CurrentState())
}
