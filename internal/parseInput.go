package statistics

import (
	"fmt"
	"strconv"
)

func ParseInput(args []string) ([]int, []string, error) {
	nums := make([]int, 0, len(args))
	warnings := make([]string, 0)

	for _, arg := range args {
		val, err := strconv.Atoi(arg)
		if err != nil {
			warnings = append(warnings, fmt.Sprintf("Skipping invalid number: %s", arg))
			continue
		}
		nums = append(nums, val)
	}

	if len(nums) == 0 {
		return nil, warnings, fmt.Errorf("no valid numbers provided")
	}

	return nums, warnings, nil
}
