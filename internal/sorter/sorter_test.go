package sorter

import (
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name      string
		width     float64
		height    float64
		length    float64
		mass      float64
		wantStack string
		wantErr   bool
	}{
		// Standard packages
		{
			name:  "standard package",
			width: 10, height: 10, length: 10, mass: 5,
			wantStack: "STANDARD",
		},

		// Heavy packages (mass >= 20, not bulky)
		{
			name:  "heavy package exactly 20kg",
			width: 10, height: 10, length: 10, mass: 20,
			wantStack: "SPECIAL",
		},
		{
			name:  "heavy package over 20kg",
			width: 10, height: 10, length: 10, mass: 25,
			wantStack: "SPECIAL",
		},

		// Bulky packages (volume >= 1,000,000, not heavy)
		{
			name:  "bulky package by volume exactly 1,000,000 cm³",
			width: 100, height: 100, length: 100, mass: 5,
			wantStack: "SPECIAL",
		},
		{
			name:  "bulky package by volume over 1,000,000 cm³",
			width: 200, height: 100, length: 100, mass: 5,
			wantStack: "SPECIAL",
		},
		{
			name:  "bulky package by dimension exactly 150cm wide",
			width: 150, height: 10, length: 10, mass: 5,
			wantStack: "SPECIAL",
		},
		{
			name:  "bulky package by dimension exactly 150cm tall",
			width: 10, height: 150, length: 10, mass: 5,
			wantStack: "SPECIAL",
		},
		{
			name:  "bulky package by dimension exactly 150cm long",
			width: 10, height: 10, length: 150, mass: 5,
			wantStack: "SPECIAL",
		},
		{
			name:  "bulky package by dimension over 150cm",
			width: 200, height: 10, length: 10, mass: 5,
			wantStack: "SPECIAL",
		},

		// Rejected packages (both bulky and heavy)
		{
			name:  "rejected package - heavy and bulky by volume",
			width: 100, height: 100, length: 100, mass: 20,
			wantStack: "REJECTED",
		},
		{
			name:  "rejected package - heavy and bulky by dimension",
			width: 150, height: 10, length: 10, mass: 20,
			wantStack: "REJECTED",
		},

		// Boundary: just under bulky/heavy thresholds → STANDARD
		{
			name:  "just under volume threshold",
			width: 99, height: 100, length: 100, mass: 5,
			wantStack: "STANDARD",
		},
		{
			name:  "just under dimension threshold",
			width: 149, height: 10, length: 10, mass: 5,
			wantStack: "STANDARD",
		},
		{
			name:  "just under heavy threshold",
			width: 10, height: 10, length: 10, mass: 19,
			wantStack: "STANDARD",
		},
		{
			name:  "just under both thresholds",
			width: 149, height: 10, length: 10, mass: 19,
			wantStack: "STANDARD",
		},

		// Invalid input
		{
			name:  "zero width",
			width: 0, height: 10, length: 10, mass: 10,
			wantErr: true,
		},
		{
			name:  "negative height",
			width: 10, height: -1, length: 10, mass: 10,
			wantErr: true,
		},
		{
			name:  "all fields invalid",
			width: 0, height: 0, length: 0, mass: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sort(tt.width, tt.height, tt.length, tt.mass)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Sort() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.wantStack {
				t.Errorf("Sort() = %q, want %q", got, tt.wantStack)
			}
		})
	}
}
