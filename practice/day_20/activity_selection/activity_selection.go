package activityselection

import "sort"

// Activity represents an activity with start and finish times
type Activity struct {
	Start  int
	Finish int
}

// SelectActivities selects the maximum number of non-overlapping activities
func SelectActivities(activities []Activity) []Activity {
	// Sort the activity by the end time
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].Finish < activities[j].Finish
	})

	result := make([]Activity, 0)
	lastEnd := 0

	for i := 0; i < len(activities); i++ {
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
