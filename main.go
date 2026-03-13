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

	lines, err := statistics.ReadInputLines(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	nums, warnings, err := statistics.ParseInput(lines)
	for _, warning := range warnings {
		fmt.Fprintln(os.Stderr, warning)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	summary := statistics.CalculateSummary(nums)

	fmt.Println("Average:", summary.Average)
	fmt.Println("Median:", summary.Median)
	fmt.Println("Variance:", summary.Variance)
	fmt.Println("Standard Deviation:", summary.StandardDeviation)
}
