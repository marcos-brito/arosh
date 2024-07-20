package history

import (
	"os"
	"testing"
)

func TestGettingLineContent(t *testing.T) {
	tests := []struct {
		content  string
		line     int
		expected string
	}{
		{
			"echo 123\nls\nmkdir hello\n",
			2,
			"mkdir hello",
		},
		{
			"docker container\n",
			0,
			"docker container",
		},
	}

	for _, tt := range tests {
		file, err := os.CreateTemp("", "")
		file.WriteString(tt.content)

		if err != nil {
			t.Errorf("Error creating temporary files")
		}

		history := NewHistory(file.Name())
		history.currentLine = tt.line

		if tt.expected != history.findLineContent() {
			t.Errorf("Expected content to be %s, but got %s", tt.expected, history.findLineContent())
		}

		defer file.Close()
	}
}

func TestReadingTotalOfLines(t *testing.T) {
	tests := []struct {
		content  string
		expected int
	}{
		{
			"echo 123\nls\nmkdir hello\n",
			3,
		},
		{
			"docker container\n",
			1,
		},
		{
			"docker container\nnpm run dev\ngo run main.go",
			3,
		},
	}

	for _, tt := range tests {
		file, err := os.CreateTemp("", "")
		file.WriteString(tt.content)

		if err != nil {
			t.Errorf("Error creating temporary files: %s", err)
		}

		history := NewHistory(file.Name())

		totalOfLines, err := history.readTotalOfLines()

		if err != nil {
			t.Errorf("Error reading total of lines: %s", err)
		}

		if tt.expected != totalOfLines {
			t.Errorf("Expected content to be %d, but got %d", tt.expected, totalOfLines)
		}

		defer file.Close()
	}

}
