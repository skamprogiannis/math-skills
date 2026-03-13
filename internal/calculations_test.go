package statistics

import "testing"

func TestCalculateSummaryRoundsValues(t *testing.T) {
	t.Parallel()

	got := CalculateSummary([]int{1, 2, 3, 4})
	want := Summary{
		Average:           3,
		Median:            3,
		Variance:          1,
		StandardDeviation: 1,
	}

	if got != want {
		t.Fatalf("summary mismatch: got %+v want %+v", got, want)
	}
}

func TestCalculateSummaryRoundsHalfUp(t *testing.T) {
	t.Parallel()

	got := CalculateSummary([]int{1, 2})
	want := Summary{
		Average:           2,
		Median:            2,
		Variance:          0,
		StandardDeviation: 1,
	}

	if got != want {
		t.Fatalf("half-up rounding mismatch: got %+v want %+v", got, want)
	}
}

func TestCalculationWrappersMatchSummary(t *testing.T) {
	t.Parallel()

	nums := []int{9, 1, 5, 3}
	summary := CalculateSummary(nums)

	if CalculateAverage(nums) != summary.Average {
		t.Fatalf("average wrapper mismatch: got %d want %d", CalculateAverage(nums), summary.Average)
	}
	if CalculateMedian(nums) != summary.Median {
		t.Fatalf("median wrapper mismatch: got %d want %d", CalculateMedian(nums), summary.Median)
	}
	if CalculateVariance(nums) != summary.Variance {
		t.Fatalf("variance wrapper mismatch: got %d want %d", CalculateVariance(nums), summary.Variance)
	}
	if CalculateStandardDeviation(nums) != summary.StandardDeviation {
		t.Fatalf("stddev wrapper mismatch: got %d want %d", CalculateStandardDeviation(nums), summary.StandardDeviation)
	}
}
