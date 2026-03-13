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

	lines, err := statistics.ReadDataLines(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	nums, parseWarnings, err := statistics.ParseInput(lines)
	for _, warning := range parseWarnings {
		fmt.Fprintln(os.Stderr, warning)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	stats := statistics.CalculateSummary(nums)

	fmt.Println("Average:", stats.Average)
	fmt.Println("Median:", stats.Median)
	fmt.Println("Variance:", stats.Variance)
	fmt.Println("Standard Deviation:", stats.StandardDeviation)
}
