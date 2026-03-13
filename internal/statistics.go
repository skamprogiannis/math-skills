package statistics

import (
	"math"
	"slices"
)

func CalculateAverage(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum / len(nums)
}

func CalculateMedian(nums []int) int {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2
	}

	return sorted[mid]
}

func CalculateVariance(nums []int) int {
	mean := CalculateAverage(nums)
	squaredDeviationSum := 0

	for _, n := range nums {
		deviation := n - mean
		squaredDeviationSum += deviation * deviation
	}

	return squaredDeviationSum / len(nums)
}

func CalculateStandardDeviation(nums []int) int {
	variance := CalculateVariance(nums)
	return int(math.Sqrt(float64(variance)))
}
