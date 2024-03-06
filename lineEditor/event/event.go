package event

type Event int
type Listener func()

const (
	LINE_ACCEPTED = iota
)

type EventManager struct {
	listeners map[Event][]Listener
}

func NewEventManager() *EventManager {
	return &EventManager{
		listeners: map[Event][]Listener{
			LINE_ACCEPTED: {},
		},
	}
}

func (eventManager *EventManager) AddListener(event Event, listener Listener) {
	eventManager.listeners[event] = append(eventManager.listeners[event], listener)
}

// TODO: I don't know how to find the function in the array. Once in, it's forever 🥲
func (eventManager *EventManager) removeListener(event Event, listener Listener) error {

	return nil
}

func (eventManager *EventManager) Notify(event Event) {
	for _, listener := range eventManager.listeners[event] {
		listener()
	}
}
