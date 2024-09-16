package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDeleteBehind(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			"",
			"",
		},
		{
			"text",
			"tex",
		},
		{
			"a",
			"",
		},
	}

	for _, tt := range tests {
		editMode := NewEmacsMode()
		shell := NewShell()
		shell.AddPlugin(editMode)
		shell.setupPlugins()
		shell.editor.SetLine(tt.input)

		editMode.deleteBehind()
		got := shell.editor.Line()
		if tt.want != got {
			t.Errorf(cmp.Diff(tt.want, got))
		}

	}

}

func TestDeleteAllBehind(t *testing.T) {
	tests := []struct {
		input    string
		position int
		want     string
	}{
		{
			"sudo pacman -Scc",
			7,
			"cman -Scc",
		},
		{
			"git commit --amend",
			18,
			"",
		},
		{
			"ls -la",
			0,
			"ls -la",
		},
	}

	for _, tt := range tests {
		editMode := NewEmacsMode()
		shell := NewShell()
		shell.AddPlugin(editMode)
		shell.setupPlugins()
		shell.editor.SetLine(tt.input)
		shell.editor.SetPostion(tt.position)

		editMode.deleteAllBehind()
		got := shell.editor.Line()
		if tt.want != got {
			t.Errorf("\nwant: %s\ngot: %s", tt.want, got)
		}

		if shell.editor.Position() != 0 {
			t.Errorf("Expected postion to be 0")
		}

	}

}

func TestDeleteAllAhead(t *testing.T) {
	tests := []struct {
		input    string
		position int
		want     string
	}{
		{
			"sudo pacman -Scc",
			7,
			"sudo pa",
		},
		{
			"git commit --amend",
			18,
			"git commit --amend",
		},
		{
			"ls -la",
			0,
			"",
		},
	}

	for _, tt := range tests {
		editMode := NewEmacsMode()
		shell := NewShell()
		shell.AddPlugin(editMode)
		shell.setupPlugins()
		shell.editor.SetLine(tt.input)
		shell.editor.SetPostion(tt.position)

		editMode.deleteAllAhead()
		got := shell.editor.Line()
		if tt.want != got {
			t.Errorf("\nwant: %s\ngot: %s", tt.want, got)
		}

		if shell.editor.Position() != tt.position {
			t.Errorf("Expected postion to be %d", tt.position)
		}

	}

}
