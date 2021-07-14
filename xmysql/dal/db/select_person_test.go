package db

import "testing"

func TestSelectPerson(t *testing.T) {
	Init()
	tests := []struct {
		name string
	}{
		{"cindy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectPerson()
		})
	}
}

func TestChildSelectPerson(t *testing.T) {
	Init()
	tests := []struct {
		name string
	}{
		{"alan"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ChildSelectPerson()
		})
	}
}
