package event

import "fmt"

type Listener func()

type EventManager struct {
	events map[string][]Listener
}

func New() *EventManager {
	return &EventManager{
		events: map[string][]Listener{},
	}
}

func (e *EventManager) add(event string) {
	_, ok := e.events[event]

	if ok {
		return
	}

	e.events[event] = []Listener{}

}

func (e *EventManager) Listen(event string, listener Listener) error {
	_, ok := e.events[event]

	if ok {
		e.events[event] = append(e.events[event], listener)
		return nil
	}

	return fmt.Errorf("Event %s not found", event)
}

func (e *EventManager) Emit(event string) {
	e.add(event)
	for _, listener := range e.events[event] {
		listener()
	}
}
