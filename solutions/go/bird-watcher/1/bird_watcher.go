package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
// TotalBirdCount returns the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    total := 0
    for _, count := range birdsPerDay {
        total += count
    }
    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    start := (week - 1) * 7  // Adjust for 1-based week indexing
    end := start + 7
    if end > len(birdsPerDay) {
        end = len(birdsPerDay)
    }
    
    total := 0
    for i := start; i < end; i++ {
        total += birdsPerDay[i]
    }
    return total
}
// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i++ {
        if i%2 == 0 {  // Changed from i%2 != 0 to i%2 == 0
            birdsPerDay[i]++  // Increment even-indexed days by 1
        }
    }
    return birdsPerDay
}