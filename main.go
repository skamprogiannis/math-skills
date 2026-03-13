package main

import (
	"fmt"
	"os"

	statistics "math-skills/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: No data file provided")
		os.Exit(1)
	}

	nums, warnings, err := statistics.ParseInput(os.Args[1:])
	for _, warning := range warnings {
		fmt.Fprintln(os.Stderr, warning)	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	average := statistics.CalculateAverage(nums)
	median := statistics.CalculateMedian(nums)
	variance := statistics.CalculateVariance(nums)
	standardDeviation := statistics.CalculateStandardDeviation(nums)

	fmt.Printf("Average: %d\n", average)
	fmt.Printf("Median: %d\n", median)
	fmt.Printf("Variance: %d\n", variance)
	fmt.Printf("Standard Deviation: %d\n", standardDeviation)
}
