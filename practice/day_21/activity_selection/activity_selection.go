package activityselection

import "sort"

// Activity represents an activity with start and finish times
type Activity struct {
	Start  int
	Finish int
}

// SelectActivities selects the maximum number of non-overlapping activities
func SelectActivities(activities []Activity) []Activity {
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].Finish < activities[j].Finish
	})

	result := []Activity{}
	lastEnd := 0

	for i := range activities {
		activity := activities[i]
		if activity.Start >= lastEnd {
			result = append(result, activity)
			lastEnd = activity.Finish
		}
	}

	return result
}

// MaxActivities returns the maximum number of non-overlapping activities
func MaxActivities(activities []Activity) int {
	return len(SelectActivities(activities))
}
