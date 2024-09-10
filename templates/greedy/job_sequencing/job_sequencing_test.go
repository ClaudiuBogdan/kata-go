package jobsequencing

import (
	"reflect"
	"testing"
)

func TestScheduleJobs(t *testing.T) {
	tests := []struct {
		name           string
		jobs           []Job
		expectedOrder  []int
		expectedProfit int
	}{
		{
			name: "Simple case",
			jobs: []Job{
				{ID: 1, Deadline: 2, Profit: 100},
				{ID: 2, Deadline: 1, Profit: 19},
				{ID: 3, Deadline: 2, Profit: 27},
				{ID: 4, Deadline: 1, Profit: 25},
				{ID: 5, Deadline: 3, Profit: 15},
			},
			expectedOrder:  []int{1, 3, 5},
			expectedProfit: 142,
		},
		{
			name: "All jobs can be scheduled",
			jobs: []Job{
				{ID: 1, Deadline: 1, Profit: 30},
				{ID: 2, Deadline: 2, Profit: 40},
				{ID: 3, Deadline: 3, Profit: 50},
			},
			expectedOrder:  []int{1, 2, 3},
			expectedProfit: 120,
		},
		{
			name: "Some jobs cannot be scheduled",
			jobs: []Job{
				{ID: 1, Deadline: 1, Profit: 30},
				{ID: 2, Deadline: 1, Profit: 40},
				{ID: 3, Deadline: 1, Profit: 50},
			},
			expectedOrder:  []int{3},
			expectedProfit: 50,
		},
		{
			name:           "Empty input",
			jobs:           []Job{},
			expectedOrder:  []int{},
			expectedProfit: 0,
		},
		{
			name: "Single job",
			jobs: []Job{
				{ID: 1, Deadline: 1, Profit: 100},
			},
			expectedOrder:  []int{1},
			expectedProfit: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schedule := ScheduleJobs(tt.jobs)
			if !reflect.DeepEqual(schedule, tt.expectedOrder) {
				t.Errorf("ScheduleJobs() = %v, want %v", schedule, tt.expectedOrder)
			}

			profit := CalculateTotalProfit(tt.jobs, schedule)
			if profit != tt.expectedProfit {
				t.Errorf("CalculateTotalProfit() = %v, want %v", profit, tt.expectedProfit)
			}

			if !isValidSchedule(tt.jobs, schedule) {
				t.Errorf("Invalid schedule: %v", schedule)
			}
		})
	}
}

func isValidSchedule(jobs []Job, schedule []int) bool {
	deadlines := make(map[int]int)
	for _, job := range jobs {
		deadlines[job.ID] = job.Deadline
	}

	for i, jobID := range schedule {
		if deadlines[jobID] < i+1 {
			return false
		}
	}
	return true
}
