package statistics

import "slices"

func calculateAverage(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum / len(nums)
}

func calculateMedian(nums []int) int {
	slices.Sort(nums)
	return nums[len(nums)/2]
}

func calculateVariance(nums []int) int {
	squaredDeviationSum := 0
	mean := calculateAverage(nums)

	for _, n := range nums {
		deviation := n - mean
		squaredDeviationSum += deviation * deviation
	}
	return deviationSum / len(nums)
}

func calculateStandardDeviation(nums []int) int {
	variance := calculateVariance(nums)
	return variance * variance
}

