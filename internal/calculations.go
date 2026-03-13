package statistics

import (
	"math"
	"slices"
)

type Summary struct {
	Average           int
	Median            int
	Variance          int
	StandardDeviation int
}

func CalculateSummary(nums []int) Summary {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	average := float64(sum) / float64(len(nums))

	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	mid := len(sorted) / 2
	median := float64(sorted[mid])
	if len(sorted)%2 == 0 {
		median = float64(sorted[mid-1]+sorted[mid]) / 2
	}

	squaredDeviationSum := 0.0
	for _, n := range nums {
		deviation := float64(n) - average
		squaredDeviationSum += deviation * deviation
	}
	variance := squaredDeviationSum / float64(len(nums))

	return Summary{
		Average:           int(math.Round(average)),
		Median:            int(math.Round(median)),
		Variance:          int(math.Round(variance)),
		StandardDeviation: int(math.Round(math.Sqrt(variance))),
	}
}

func CalculateAverage(nums []int) int {
	return CalculateSummary(nums).Average
}

func CalculateMedian(nums []int) int {
	return CalculateSummary(nums).Median
}

func CalculateVariance(nums []int) int {
	return CalculateSummary(nums).Variance
}

func CalculateStandardDeviation(nums []int) int {
	return CalculateSummary(nums).StandardDeviation
}
