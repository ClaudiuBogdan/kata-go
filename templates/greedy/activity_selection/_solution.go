package activityselection

import (
	"sort"
)

// Activity represents an activity with start and finish times
type Activity struct {
	Start  int
	Finish int
}

// SelectActivities selects the maximum number of non-overlapping activities
func SelectActivities(activities []Activity) []Activity {
	if len(activities) == 0 {
		return []Activity{}
	}

	// Sort activities by finish time
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].Finish < activities[j].Finish
	})

	selected := []Activity{activities[0]}
	lastFinish := activities[0].Finish

	for i := 1; i < len(activities); i++ {
		if activities[i].Start >= lastFinish {
			selected = append(selected, activities[i])
			lastFinish = activities[i].Finish
		}
	}

	return selected
}

// MaxActivities returns the maximum number of non-overlapping activities
func MaxActivities(activities []Activity) int {
	return len(SelectActivities(activities))
}
