package event

import (
	"testing"
)

// WARN: I didn't used much time to think on how to test this. But the module is
// quite dumb so maybe it's enough
func TestNotifying(t *testing.T) {
	tests := []struct {
		event    Event
		expected string
	}{
		{

			LINE_ACCEPTED,
			"this is changing",
		},
		{
			LINE_ACCEPTED,
			"chaging as well",
		},
	}

	for _, tt := range tests {
		eventManager := NewEventManager()
		toBeChanged := ""

		eventManager.AddListener(LINE_ACCEPTED, func() {
			toBeChanged = tt.expected
		})
		eventManager.Notify(LINE_ACCEPTED)

		if toBeChanged != tt.expected {
			t.Errorf("Expected \"%s\" to be \"%s\" after event", toBeChanged, tt.expected)
		}
	}
}
