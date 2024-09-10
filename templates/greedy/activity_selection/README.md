# Greedy Algorithm for Activity Selection Problem

## Problem Description

The Activity Selection Problem involves selecting the maximum number of non-overlapping activities given their start and finish times. Each activity is represented by a start time and a finish time, and the goal is to select as many activities as possible such that no two selected activities overlap.

## Approach

This implementation uses a greedy algorithm to solve the Activity Selection Problem. The algorithm works as follows:

1. Sort all activities in ascending order of finish time.
2. Select the first activity (the one that finishes earliest).
3. For the remaining activities, select an activity if its start time is greater than or equal to the finish time of the previously selected activity.

This greedy choice of always selecting the next possible activity that finishes earliest leads to an optimal solution.

## Complexity Analysis

- Time Complexity: O(n log n), where n is the number of activities. This is due to the sorting step.
- Space Complexity: O(n) to store the sorted activities and the selected activities.

## Usage

```go
activities := []Activity{
    {1, 4}, {3, 5}, {0, 6}, {5, 7}, {3, 9}, {5, 9}, {6, 10}, {8, 11}, {8, 12}, {2, 14}, {12, 16},
}

selected := SelectActivities(activities)
fmt.Println("Selected activities:", selected)

maxCount := MaxActivities(activities)
fmt.Println("Maximum number of activities:", maxCount)
```

## Implementation Details

The package provides the following main components:

1. `Activity` struct: Represents an activity with start and finish times.
2. `SelectActivities(activities []Activity) []Activity`: Selects the maximum number of non-overlapping activities.
3. `MaxActivities(activities []Activity) int`: Returns the maximum number of non-overlapping activities.

The implementation uses Go's built-in `sort` package to sort the activities by finish time.

## Testing

The implementation includes a test suite that covers various scenarios:

1. Simple case with a mix of overlapping and non-overlapping activities
2. All non-overlapping activities
3. All overlapping activities
4. Empty input
5. Single activity

To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of the activity selection implementation under different scenarios.

## Advantages and Limitations

Advantages:
- Efficiently selects the maximum number of non-overlapping activities
- Simple to understand and implement
- Optimal solution guaranteed for this problem

Limitations:
- Assumes that activities can be sorted by finish time
- May not be suitable for more complex scheduling problems with additional constraints

## Applications

- Scheduling problems in various domains (e.g., job scheduling, resource allocation)
- Time management and planning
- Room booking systems
- Transportation and logistics (e.g., scheduling delivery time slots)
- Network packet scheduling

Remember that while this greedy approach works optimally for the basic Activity Selection Problem, more complex scheduling problems may require different algorithms or additional considerations.

