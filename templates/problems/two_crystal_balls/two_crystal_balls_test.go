package twocrystalballs

import "testing"

func TestTwoCrystalBalls(t *testing.T) {
	tests := []struct {
		name     string
		breaks   []bool
		expected int
	}{
		{
			name:     "Break at middle",
			breaks:   []bool{false, false, false, false, true, true, true, true},
			expected: 4,
		},
		{
			name:     "Break at first",
			breaks:   []bool{true, true, true, true},
			expected: 0,
		},
		{
			name:     "Break at last",
			breaks:   []bool{false, false, false, true},
			expected: 3,
		},
		{
			name:     "No break",
			breaks:   []bool{false, false, false, false},
			expected: -1,
		},
		{
			name:     "Large input",
			breaks:   append(make([]bool, 999), true),
			expected: 999,
		},
		{
			name:     "Empty input",
			breaks:   []bool{},
			expected: -1,
		},
		{
			name:     "Single element (true)",
			breaks:   []bool{true},
			expected: 0,
		},
		{
			name:     "Single element (false)",
			breaks:   []bool{false},
			expected: -1,
		},
		{
			name:     "Two elements (break at second)",
			breaks:   []bool{false, true},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TwoCrystalBalls(tt.breaks)
			if result != tt.expected {
				t.Errorf("TwoCrystalBalls() = %v, want %v", result, tt.expected)
			}
		})
	}
}
