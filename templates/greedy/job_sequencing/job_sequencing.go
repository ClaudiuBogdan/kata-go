package jobsequencing

import (
	"sort"
)

// Job represents a job with its deadline and profit
type Job struct {
	ID       int
	Deadline int
	Profit   int
}

// ScheduleJobs schedules jobs to maximize profit
func ScheduleJobs(jobs []Job) []int {
	n := len(jobs)
	if n == 0 {
		return []int{}
	}

	// Sort jobs in decreasing order of profit
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Profit > jobs[j].Profit
	})

	// Find the maximum deadline
	maxDeadline := 0
	for _, job := range jobs {
		if job.Deadline > maxDeadline {
			maxDeadline = job.Deadline
		}
	}

	// Initialize the schedule
	schedule := make([]int, maxDeadline)
	for i := range schedule {
		schedule[i] = -1 // -1 indicates an empty slot
	}

	// Schedule jobs
	for _, job := range jobs {
		for i := job.Deadline - 1; i >= 0; i-- {
			if schedule[i] == -1 {
				schedule[i] = job.ID
				break
			}
		}
	}

	// Remove empty slots from the schedule
	result := []int{}
	for _, jobID := range schedule {
		if jobID != -1 {
			result = append(result, jobID)
		}
	}

	return result
}

// CalculateTotalProfit calculates the total profit of the scheduled jobs
func CalculateTotalProfit(jobs []Job, schedule []int) int {
	totalProfit := 0
	jobMap := make(map[int]int)
	for _, job := range jobs {
		jobMap[job.ID] = job.Profit
	}
	for _, jobID := range schedule {
		totalProfit += jobMap[jobID]
	}
	return totalProfit
}
