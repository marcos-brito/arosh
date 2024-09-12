package event

import (
	"testing"
)

func TestNotifying(t *testing.T) {
	tests := []struct {
		event    string
		expected string
	}{
		{

			"EVENT_1",
			"this is changing",
		},
	}

	for _, tt := range tests {
		eventManager := New()
		toBeChanged := ""

		eventManager.Listen(tt.event, func() {
			toBeChanged = tt.expected
		})

		eventManager.Emit(tt.event)

		if toBeChanged != tt.expected {
			t.Errorf("Expected \"%s\" to be \"%s\" after event", toBeChanged, tt.expected)
		}
	}
}
