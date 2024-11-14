package activityselection

import "sort"

// Activity represents an activity with start and finish times
type Activity struct {
	Start  int
	Finish int
}

// SelectActivities selects the maximum number of non-overlapping activities
func SelectActivities(activities []Activity) []Activity {
	result := make([]Activity, 0)
	// Sort activities by finish
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].Finish < activities[j].Finish
	})

	lastEnd := 0

	for i := range activities {
		activity := activities[i]
		if activity.Start >= lastEnd {
			lastEnd = activity.Finish
			result = append(result, activity)
		}
	}

	return result
}

// MaxActivities returns the maximum number of non-overlapping activities
func MaxActivities(activities []Activity) int {
	result := SelectActivities(activities)

	return len(result)
}
