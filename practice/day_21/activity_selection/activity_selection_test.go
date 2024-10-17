package activityselection

import (
	"reflect"
	"testing"
)

func TestSelectActivities(t *testing.T) {
	tests := []struct {
		name     string
		activities []Activity
		expected []Activity
	}{
		{
			name: "Simple case",
			activities: []Activity{
				{1, 4}, {3, 5}, {0, 6}, {5, 7}, {3, 9}, {5, 9}, {6, 10}, {8, 11}, {8, 12}, {2, 14}, {12, 16},
			},
			expected: []Activity{{1, 4}, {5, 7}, {8, 11}, {12, 16}},
		},
		{
			name: "All non-overlapping",
			activities: []Activity{
				{1, 2}, {3, 4}, {5, 6}, {7, 8},
			},
			expected: []Activity{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
		},
		{
			name: "All overlapping",
			activities: []Activity{
				{1, 5}, {2, 6}, {3, 7}, {4, 8},
			},
			expected: []Activity{{1, 5}},
		},
		{
			name:     "Empty input",
			activities: []Activity{},
			expected: []Activity{},
		},
		{
			name: "Single activity",
			activities: []Activity{{1, 2}},
			expected: []Activity{{1, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SelectActivities(tt.activities)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SelectActivities() = %v, want %v", result, tt.expected)
			}

			if len(result) != MaxActivities(tt.activities) {
				t.Errorf("MaxActivities() = %d, want %d", MaxActivities(tt.activities), len(result))
			}

			if !isValidSelection(tt.activities, result) {
				t.Errorf("Invalid selection: %v", result)
			}
		})
	}
}

func isValidSelection(all, selected []Activity) bool {
	if len(selected) == 0 {
		return true
	}

	// Check if selected activities are in the original set
	selectedMap := make(map[Activity]bool)
	for _, a := range selected {
		selectedMap[a] = true
	}
	for _, a := range all {
		if selectedMap[a] {
			delete(selectedMap, a)
		}
	}
	if len(selectedMap) > 0 {
		return false
	}

	// Check for non-overlapping
	for i := 0; i < len(selected)-1; i++ {
		if selected[i].Finish > selected[i+1].Start {
			return false
		}
	}

	return true
}
